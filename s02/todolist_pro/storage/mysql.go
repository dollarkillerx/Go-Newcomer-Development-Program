package storage

import (
	"github.com/dollarkillerx/Go-Newcomer-Development-Program/s02/todolist_pro/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"errors"
	"fmt"
	"log"
)

// 实现基于MySQL的TodoListStorage   学习MySQL & GORM https://gorm.io/
type MySQLStorage struct {
	db *gorm.DB
}

// 定义GORM映射模型
type BankAccount struct {
	gorm.Model
	Account  string `gorm:"index"` // 加上索引
	Password string `gorm:"index"`
	Balance  float32
}

func NewMySQLAccount() TodoListStorage {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Config.MySQLConf.User, conf.Config.MySQLConf.Password, conf.Config.MySQLConf.Address, conf.Config.MySQLConf.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// 自动建表
	if err := db.AutoMigrate(&BankAccount{}); err != nil {
		log.Fatalln(err)
	}
	return &MySQLStorage{
		db: db,
	}
}

func (m *MySQLStorage) OpenAccount(account, password string) error {
	act := &BankAccount{
		Account:  account,
		Password: account,
	}

	return m.db.Model(&BankAccount{}).Create(&act).Error
}

func (m *MySQLStorage) CheckAccount(account, password string) error {
	var act BankAccount
	if err := m.db.Model(&BankAccount{}).
		Where("account =?", account).
		Where("password =?", password).
		Find(&act).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 检测错误, ErrRecordNotFound是没有找到
			return err
		}
		return err
	}

	fmt.Println(act)
	return nil
}

func (m *MySQLStorage) BalanceEnquiry(account string) (balance float32, err error) {
	var act BankAccount
	if err := m.db.Model(&BankAccount{}).
		Where("account =?", account).
		Find(&act).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 检测错误, ErrRecordNotFound是没有找到
			return 0, err
		}
		return 0, err
	}

	return act.Balance, nil
}

func (m *MySQLStorage) Withdrawals(account string, money float32) error {
	var act BankAccount
	if err := m.db.Model(&BankAccount{}).
		Where("account =?", account).
		Find(&act).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 检测错误, ErrRecordNotFound是没有找到
			return err
		}
		return err
	}

	if act.Balance < money {
		return errors.New("余额不足")
	}

	act.Balance += money
	if err := m.db.Model(&BankAccount{}).
		Where("account =?", account).
		Updates(act).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 检测错误, ErrRecordNotFound是没有找到
			return err
		}
		return err
	}

	return nil
}

func (m *MySQLStorage) Deposit(account string, money float32) error {
	var act BankAccount
	if err := m.db.Model(&BankAccount{}).
		Where("account =?", account).
		Find(&act).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 检测错误, ErrRecordNotFound是没有找到
			return err
		}
		return err
	}

	act.Balance += money
	if err := m.db.Model(&BankAccount{}).
		Where("account =?", account).
		Updates(act).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 检测错误, ErrRecordNotFound是没有找到
			return err
		}
		return err
	}

	return nil
}
