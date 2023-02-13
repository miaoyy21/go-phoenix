package handle

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/base"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Context struct {
	*sql.DB
	*http.Request

	Writer http.ResponseWriter
	params map[string]string
	values map[string]string

	token      string
	userId     string
	userCode   string
	userName   string
	departId   string
	departCode string
	departName string
}

func NewContext(db *sql.DB, r *http.Request, w http.ResponseWriter) *Context {
	return &Context{DB: db, Request: r, Writer: w}
}

func (ctx *Context) HasInWhiteRoute(method string, path string, params map[string]string, values map[string]string) bool {
	route, ok := whiteRoutes[method][path]
	if !ok {
		return false
	}

	// Params
	for key, value := range route.Params {
		val, ok := params[key]
		if !ok {
			return false
		}

		if !strings.EqualFold(val, value) {
			return false
		}
	}

	// Values
	for key, value := range route.Values {
		val, ok := values[key]
		if !ok {
			return false
		}

		if !strings.EqualFold(val, value) {
			return false
		}
	}

	return true
}

func (ctx *Context) Parse() error {

	// Token
	token, err := ctx.Cookie("PHOENIX_LOGIN_TOKEN")
	if err != nil {
		return err
	}
	ctx.token = token.Value

	// User ID
	userId, err := ctx.Cookie("PHOENIX_USER_ID")
	if err != nil {
		return err
	}
	ctx.userId = userId.Value

	// User code
	userCode, err := ctx.Cookie("PHOENIX_USER_CODE")
	if err != nil {
		return err
	}
	ctx.userCode = userCode.Value

	// User name
	userName, err := ctx.Cookie("PHOENIX_USER_NAME")
	if err != nil {
		return err
	}

	xUserName, err := base64.StdEncoding.DecodeString(userName.Value)
	if err != nil {
		return err
	}
	ctx.userName = string(xUserName)

	// Depart ID
	departId, err := ctx.Cookie("PHOENIX_DEPART_ID")
	if err != nil {
		return err
	}
	ctx.departId = departId.Value

	// Depart code
	departCode, err := ctx.Cookie("PHOENIX_DEPART_CODE")
	if err != nil {
		return err
	}
	ctx.departCode = departCode.Value

	// Depart name
	departName, err := ctx.Cookie("PHOENIX_DEPART_NAME")
	if err != nil {
		return err
	}

	xDepartName, err := base64.StdEncoding.DecodeString(departName.Value)
	if err != nil {
		return err
	}
	ctx.departName = string(xDepartName)

	// 是否为有效的Token
	if err := ctx.parseToken(); err != nil {
		return fmt.Errorf("parse Token Failure :: %s", err.Error())
	}

	return nil
}

func (ctx *Context) parseToken() error {
	bytes, err := base64.StdEncoding.DecodeString(ctx.token)
	if err != nil {
		return err
	}

	src := strings.Split(string(base.Config.AesStream(bytes)), ",")
	if len(src) != 3 {
		return errors.New("token Split count is not Match")
	}

	// 0 是否与 用户ID 匹配
	if !strings.EqualFold(src[0], ctx.GetUserId()) {
		return errors.New("user Id is not Match")
	}

	// 查询用户的加密密码
	var userPassword string
	row := ctx.QueryRow("SELECT password_ FROM sys_user WHERE id = ?", ctx.GetUserId())
	if err := row.Scan(&userPassword); err != nil {
		return err
	}

	// 1 是否与 附加信息 匹配
	ext := fmt.Sprintf("%s_%s_%s_%s_%s_%s_%s", ctx.UserAgent(), userPassword, ctx.GetUserCode(), ctx.GetUserName(), ctx.GetDepartId(), ctx.GetDepartCode(), ctx.GetDepartName())
	md5Ext := md5.Sum([]byte(ext))
	if !strings.EqualFold(src[1], base64.StdEncoding.EncodeToString(md5Ext[:])) {
		return errors.New("user's Agent is not Match")
	}

	// 2 是否超过失效时间
	expire, err := strconv.ParseInt(src[2], 10, 64)
	if err != nil {
		return err
	}

	// Max Time Exceed
	if time.Now().After(time.Unix(expire, 0)) {
		return errors.New("token is Expired")
	}

	return nil
}

func (ctx *Context) GetUserId() string {
	return ctx.userId
}

func (ctx *Context) GetUserCode() string {
	return ctx.userCode
}

func (ctx *Context) GetUserName() string {
	return ctx.userName
}

func (ctx *Context) GetDepartId() string {
	return ctx.departId
}

func (ctx *Context) GetDepartCode() string {
	return ctx.departCode
}

func (ctx *Context) GetDepartName() string {
	return ctx.departName
}

func (ctx *Context) GetNullableFormValue(key string) interface{} {
	sParent := ctx.PostFormValue(key)
	if len(sParent) < 1 || strings.EqualFold(sParent, "0") {
		return nil
	}

	return sParent
}

func (ctx *Context) GetSortsFilters(mapFields map[string]string) ([]string, []string) {
	sorts := make([]string, 0)
	filters := make([]string, 0)

	// 多列排序时，必须依次从最原始的请求参数中获取排序字段
	uri, err := url.ParseRequestURI(ctx.RequestURI)
	if err != nil {
		logrus.Errorf("url.ParseRequestURI(%s) Failure :: %s", ctx.RequestURI, err.Error())
		return sorts, filters
	}

	params := strings.Split(uri.RawQuery, "&")
	logrus.Debugf("Params is %#v\n", params)
	for _, param := range params {
		pairs := strings.Split(param, "=")
		if len(pairs) != 2 {
			continue
		}

		key, value := pairs[0], pairs[1]
		if len(value) < 1 {
			continue
		}
		logrus.Debugf("Params key is %s && Value is %s \n", key, value)

		// 符合排序规则 sort[...]={asc||desc}
		if strings.HasPrefix(key, "sort[") && strings.HasSuffix(key, "]") {
			col, asc := key[5:len(key)-1], strings.ToUpper(value)
			if !strings.EqualFold(asc, "ASC") && !strings.EqualFold(asc, "DESC") {
				continue
			}

			if newKey, ok := mapFields[col]; ok {
				sorts = append(sorts, fmt.Sprintf("%s %s", newKey, asc))
			} else {
				sorts = append(sorts, fmt.Sprintf("%s %s", col, asc))
			}
		}

		// 符合过滤规则 filter[...]={value}
		if strings.HasPrefix(key, "filter[") && strings.HasSuffix(key, "]") {
			col := key[7 : len(key)-1]

			val, err := url.PathUnescape(strings.ToUpper(value))
			if err != nil {
				logrus.Errorf("url.PathUnescape(%s) Failure :: %s", strings.ToUpper(value), err.Error())
				continue
			}

			// 如果为空格，那么过滤全部
			if newKey, ok := mapFields[col]; ok {
				filters = append(filters, fmt.Sprintf("UPPER(%s) LIKE '%%%s%%'", newKey, val))
			} else {
				filters = append(filters, fmt.Sprintf("UPPER(%s) LIKE '%%%s%%'", col, val))
			}
		}
	}

	return sorts, filters
}

func (ctx *Context) GetIP() string {
	ip, _, err := net.SplitHostPort(strings.TrimSpace(ctx.RemoteAddr))
	if err != nil {
		return "0.0.0.0"
	}

	remoteIP := net.ParseIP(ip)
	if remoteIP == nil {
		return "0.0.0.0"
	}

	return remoteIP.String()
}

func (ctx *Context) GetSize() string {
	return strconv.Itoa(int(ctx.ContentLength))
}

func (ctx *Context) GetMethod() string {
	return ctx.Method
}

func (ctx *Context) GetPath() string {
	return ctx.URL.Path
}

// GetUsingMenu 默认从请求地址获取当前打开的界面，否则从Cookie获取
func (ctx *Context) GetUsingMenu() string {
	// Params
	values, err := url.ParseQuery(ctx.URL.RawQuery)
	if err != nil {
		return ""
	}

	menu := values.Get("PHOENIX_USING_MENU")
	if len(menu) < 1 {
		cookie, err := ctx.Cookie("PHOENIX_USING_MENU")
		if err != nil {
			logrus.Errorf("Cookie(%q) ERROR : %s", "PHOENIX_USING_MENU", err.Error())
		} else {
			menu = strings.Trim(cookie.Value, "%22")
		}
	}

	return menu
}

func (ctx *Context) GetParams() map[string]string {
	if ctx.params == nil {
		values, err := url.ParseQuery(ctx.URL.RawQuery)
		if err != nil {
			return make(map[string]string)
		}

		delete(values, "PHOENIX_USING_MENU")
		ctx.params = base.GetURLValues(values)
	}

	return ctx.params
}

func (ctx *Context) GetValues() map[string]string {
	if ctx.values == nil {
		if err := ctx.ParseForm(); err != nil {
			return make(map[string]string)
		}

		ctx.values = base.GetURLValues(ctx.PostForm)
	}

	return ctx.values
}

func (ctx *Context) Reset(params map[string]string, values map[string]string) {
	ctx.params = params
	ctx.values = values
}

//
//func (cx *Context) ParsePayload(params interface{}) (map[string]interface{}, error) {
//	fields := make(map[string]interface{})
//
//	// Parse Form
//	if err := cx.ParseForm(); err != nil {
//		return nil, err
//	}
//
//	// Get Form
//	form := cx.PostForm
//	if len(form) < 0 {
//		return nil, errors.New("invalid payload form data")
//	}
//
//	v := reflect.ValueOf(params).Elem()
//	t := reflect.TypeOf(params).Elem()
//	for i := 0; i < t.NumField(); i++ {
//		tag := t.Field(i).Tag
//		payload, ok := tag.Lookup("payload")
//		if !ok {
//			logrus.Warnf("Params %q miss struct tag payload", t.Field(i).name)
//			continue
//		}
//
//		if ok := form.Has(payload); !ok {
//			logrus.Warnf("Form has no field %q", payload)
//			continue
//		}
//
//		value, err := cx.setValue(form.Get(payload), v.Field(i))
//		if err != nil {
//			return nil, err
//		}
//
//		column, ok := tag.Lookup("column")
//		if !ok {
//			logrus.Warnf("Params %q miss struct tag column", t.Field(i).name)
//			continue
//		}
//
//		fields[column] = value
//	}
//
//	return fields, nil
//}
//
//func (cx *Context) setValue(src string, value reflect.Value) (interface{}, error) {
//	kind := value.Kind()
//	if kind == reflect.String {
//		value.SetString(src)
//		return src, nil
//	} else if kind == reflect.Int {
//		num, err := strconv.ParseInt(src, 10, 64)
//		if err != nil {
//			return nil, err
//		}
//
//		value.SetInt(num)
//		return num, nil
//	} else if kind == reflect.Float64 {
//		num, err := strconv.ParseFloat(src, 64)
//		if err != nil {
//			return nil, err
//		}
//
//		value.SetFloat(num)
//		return num, nil
//	} else if kind == reflect.Struct {
//		//_, ok := value.Interface().(time.Time)
//		//if !ok {
//		//	return nil, errors.New("only support time.Time field")
//		//}
//		//
//		//dt, err := time.Parse("2006-01-02 15:04:05", src)
//		//if err != nil {
//		//	return nil, err
//		//}
//		//
//		//value.Set(reflect.ValueOf(dt))
//		//return dt, nil
//	}
//
//	return nil, fmt.Errorf("unsupport kind type %s", kind)
//}
