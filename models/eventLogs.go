package models

type EventLogs struct {
	Id               int64   `json:"id" gorm:"type: bigint(64); not null;primaryKey;autoIncrement"`
	Block_Number     int64   `json:"block_Number" gorm:"type: bigint(64); not null"`
	Transaction_Hash string  `json:"transaction_Hash" gorm:"type: varchar(100);not null"`
	Value            string  `json:"value" gorm:"type: varchar(100);not null"`
	FromAddress      string  `json:"from_Address" gorm:"type: varchar(100);not null"`
	ToAddress        string  `json:"to_Address" gorm:"type: varchar(100);not null"`
	IsReceived       bool    `json:"is_received" gorm:"type:boolean;not null"`
	FromNetworkId    int     `json:"fromNetworkID" gorm:"type:int;not null"`
	ToNetworkID      int     `json:"toNetworkID" gorm:"type:int;not null"`
	Nonce            int     `json:"nonce" gorm:"type:string; not null"`
	FromNetwork      NetWork `gorm:"references:Id"`
	ToNetwork        NetWork `gorm:"references:Id"`
}
