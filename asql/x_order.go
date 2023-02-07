package asql

import (
	"database/sql"
	"fmt"
	"strings"
)

// Order （使用排序的依据必须是 order_ ASC）
func Order(tx *sql.Tx, table string, source string, target string, targetIndex string, targetParent string) error {
	if source == target || len(source) < 1 {
		return nil
	}

	if len(target) < 1 {
		if len(targetParent) < 1 {
			// 场景 => 在Table中将source移至尾部
			query := fmt.Sprintf("UPDATE %s SET order_ = ?  WHERE id = ?", table)
			return Update(tx, query, GenerateOrderId(), source)
		} else if strings.EqualFold(targetParent, "0") {
			// 移动至根结点的第1个
			query := fmt.Sprintf("UPDATE %s SET order_ = ?, parent_id_ = NULL  WHERE id = ?", table)
			return Update(tx, query, GenerateOrderId(), source)
		} else {
			var xOrdered int64

			// 目标顺序号
			query := fmt.Sprintf("SELECT MAX(order_) FROM %s WHERE parent_id_ = ?", table)
			if err := SelectRow(tx, query, targetParent).Scan(&xOrdered); err != nil {
				return err
			}
			xOrdered = xOrdered + 2048

			// 移动至结点的尾部
			query = fmt.Sprintf("UPDATE %s SET order_ = ?  WHERE id = ?", table)
			return Update(tx, query, xOrdered, source)
		}
	}

	// 结束行的排序号
	ordered, err := calculate(tx, table, target)
	if err != nil {
		return err
	}

	if len(targetParent) < 1 {
		query := fmt.Sprintf("UPDATE %s SET order_ = ? WHERE id = ?", table)

		// 将起始行的序号更新为 ordered
		if err := Update(tx, query, ordered, source); err != nil {
			return err
		}
	} else if strings.EqualFold(targetParent, "0") {
		// 将起始行的序号更新为 ordered
		query := fmt.Sprintf("UPDATE %s SET order_ = ?, parent_id_ = NULL WHERE id = ?", table)
		if err := Update(tx, query, ordered, source); err != nil {
			return err
		}
	} else {
		// 将起始行的序号更新为 ordered
		query := fmt.Sprintf("UPDATE %s SET order_ = ?, parent_id_ = ? WHERE id = ?", table)
		if err := Update(tx, query, ordered, targetParent, source); err != nil {
			return err
		}
	}

	return nil
}

// 10 -> 20 {10} 30
func calculate(tx *sql.Tx, table string, target string) (int64, error) {
	var xOrdered, tOrdered int64

	// 目标顺序号
	if err := SelectRow(tx, fmt.Sprintf("SELECT order_ FROM %s WHERE id = ?", table), target).Scan(&tOrdered); err != nil {
		return 0, err
	}

	// 寻找小于目标节点的最大顺序号
	if err := SelectRow(tx, fmt.Sprintf("SELECT CASE WHEN MAX(order_) IS NULL THEN ? ELSE MAX(order_) END FROM %s WHERE order_ < ? ", table), tOrdered-2048, tOrdered).Scan(&xOrdered); err != nil {
		return 0, err
	}

	ordered := (tOrdered + xOrdered) / 2
	if ordered < tOrdered && ordered > xOrdered {
		return ordered, nil
	}

	query := fmt.Sprintf("UPDATE %s SET order_ = order_ - 2048 WHERE order_ <= ?", table)
	if err := Update(tx, query, xOrdered); err != nil {
		return 0, err
	}

	return tOrdered + 1024, nil
}
