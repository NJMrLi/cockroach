// Code generated by "stringer -type=TxnStateEnum"; DO NOT EDIT.

package sql

import "fmt"

const _TxnStateEnum_name = "NoTxnFirstBatchOpenAbortedRestartWaitCommitWait"

var _TxnStateEnum_index = [...]uint8{0, 5, 15, 19, 26, 37, 47}

func (i TxnStateEnum) String() string {
	if i < 0 || i >= TxnStateEnum(len(_TxnStateEnum_index)-1) {
		return fmt.Sprintf("TxnStateEnum(%d)", i)
	}
	return _TxnStateEnum_name[_TxnStateEnum_index[i]:_TxnStateEnum_index[i+1]]
}
