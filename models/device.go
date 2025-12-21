package models

import (
	"time"
)

type Device struct {
	Id                        int       `xorm:"not null pk autoincr comment('主键id') INT"`
	Status                    string    `xorm:"comment('状态') VARCHAR(200)"`
	TextMessage               string    `xorm:"comment('文本消息') VARCHAR(200)"`
	DeviceNumber              string    `xorm:"comment('设备号') VARCHAR(200)"`
	MaintenanceFactory        string    `xorm:"comment('维护工厂') VARCHAR(200)"`
	EquipmentNumber           string    `xorm:"comment('设备编号') VARCHAR(200)"`
	UnitName                  string    `xorm:"comment('所属单位名称') VARCHAR(200)"`
	SystemActivationTime      string    `xorm:"comment('制度开动时间') VARCHAR(200)"`
	CumulativeOperatingTime   string    `xorm:"comment('截止2024年12月累计开动时间') VARCHAR(200)"`
	DeviceCategory            string    `xorm:"comment('设备种类（板块）') VARCHAR(200)"`
	DeviceDescription         string    `xorm:"comment('设备说明') VARCHAR(200)"`
	TechnicalObjectType       string    `xorm:"comment('技术对象类型') VARCHAR(200)"`
	TotalWeight               string    `xorm:"comment('总重量') VARCHAR(200)"`
	WeightUnit                string    `xorm:"comment('重量单位') VARCHAR(200)"`
	Dimensions                string    `xorm:"comment('大小尺寸') VARCHAR(200)"`
	AcquisitionValue          string    `xorm:"comment('购置值') VARCHAR(200)"`
	CurrencyUnit              string    `xorm:"comment('货币单位') VARCHAR(200)"`
	PurchaseDate              string    `xorm:"comment('购置日期') VARCHAR(200)"`
	StartDate                 string    `xorm:"comment('开始日期') VARCHAR(200)"`
	Manufacturer              string    `xorm:"comment('制造商') VARCHAR(200)"`
	ManufacturerCountry       string    `xorm:"comment('制造商国家') VARCHAR(200)"`
	Model                     string    `xorm:"comment('型号') VARCHAR(200)"`
	SerialNumber              string    `xorm:"comment('系列号') VARCHAR(200)"`
	SupplierWarrantyStart     string    `xorm:"comment('供应商担保开始') VARCHAR(200)"`
	SupplierWarrantyEnd       string    `xorm:"comment('供应商保修结束') VARCHAR(200)"`
	FactoryArea               string    `xorm:"comment('工厂区域') VARCHAR(200)"`
	AbcIndicator              string    `xorm:"comment('ABC标识') VARCHAR(200)"`
	CompanyCode               string    `xorm:"comment('公司代码') VARCHAR(200)"`
	BusinessScope             string    `xorm:"comment('业务范围') VARCHAR(200)"`
	AssetNumber               string    `xorm:"comment('资产编码') VARCHAR(200)"`
	CostCenter                string    `xorm:"comment('成本中心') VARCHAR(200)"`
	WbsElement                string    `xorm:"comment('WBS元素') VARCHAR(200)"`
	PlanningPlant             string    `xorm:"comment('计划工厂') VARCHAR(200)"`
	PlannerGroup              string    `xorm:"comment('计划人员组') VARCHAR(200)"`
	MaintenanceWorkCenter     string    `xorm:"comment('维护工作中心') VARCHAR(200)"`
	FunctionalLocation        string    `xorm:"comment('安装功能位置') VARCHAR(200)"`
	SuperiorEquipment         string    `xorm:"comment('上一级设备') VARCHAR(200)"`
	ConstructionType          string    `xorm:"comment('构造类型') VARCHAR(200)"`
	EquipmentUsageStatus      string    `xorm:"comment('设备使用状态') VARCHAR(200)"`
	EquipmentTechnicalStatus  string    `xorm:"comment('设备技术状态') VARCHAR(200)"`
	OriginalEquipmentNumber   string    `xorm:"comment('原设备编号') VARCHAR(200)"`
	OriginalCompanyCode       string    `xorm:"comment('原公司代码') VARCHAR(200)"`
	OriginalAssetNumber       string    `xorm:"comment('原资产编号') VARCHAR(200)"`
	OriginalEquipmentCategory string    `xorm:"comment('原设备分类') VARCHAR(200)"`
	MainEquipmentIdentifier   string    `xorm:"comment('主设备标识') VARCHAR(200)"`
	SpecialEquipment          string    `xorm:"comment('特种设备') VARCHAR(200)"`
	IsEquipmentInstalled      string    `xorm:"comment('是否安装设备') VARCHAR(200)"`
	EquipmentProperties       string    `xorm:"comment('设备属性') VARCHAR(200)"`
	SelfNumber                string    `xorm:"comment('自编号') VARCHAR(200)"`
	FactoryNumber             string    `xorm:"comment('出厂编号') VARCHAR(200)"`
	WellNumber                string    `xorm:"comment('井号') VARCHAR(200)"`
	ProductionDate            string    `xorm:"comment('出厂日期') VARCHAR(200)"`
	VehicleIdentifier         string    `xorm:"comment('车辆标识') VARCHAR(200)"`
	LicensePlateNumber        string    `xorm:"comment('车牌号') VARCHAR(200)"`
	ChassisModel              string    `xorm:"comment('底盘型号') VARCHAR(200)"`
	ChassisNumber             string    `xorm:"comment('底盘号') VARCHAR(200)"`
	EngineModel               string    `xorm:"comment('发动机型号/引擎型号') VARCHAR(200)"`
	EngineNumber              string    `xorm:"comment('发动机号') VARCHAR(200)"`
	Capacity                  string    `xorm:"comment('能力') VARCHAR(200)"`
	CapacityUnit              string    `xorm:"comment('能力单位') VARCHAR(200)"`
	TotalPower                string    `xorm:"comment('总功率') VARCHAR(200)"`
	EnergyConsumptionType     string    `xorm:"comment('能耗种类') VARCHAR(200)"`
	AcquisitionNature         string    `xorm:"comment('购置性质') VARCHAR(200)"`
	FundingChannel            string    `xorm:"comment('资金渠道') VARCHAR(200)"`
	UsefulLife                string    `xorm:"comment('效用年限') VARCHAR(200)"`
	OriginalAssetValue        string    `xorm:"comment('资产原值') VARCHAR(200)"`
	NetAssetValue             string    `xorm:"comment('资产净值') VARCHAR(200)"`
	AccumulatedDepreciation   string    `xorm:"comment('累计折旧') VARCHAR(200)"`
	AssetExtension            string    `xorm:"comment('资产延期') VARCHAR(200)"`
	ResidualValueRate         string    `xorm:"comment('残值率') VARCHAR(200)"`
	ContractNumber            string    `xorm:"comment('合同号') VARCHAR(200)"`
	PurchaseOrder             string    `xorm:"comment('采购订单') VARCHAR(200)"`
	Supplier                  string    `xorm:"comment('供应商') VARCHAR(200)"`
	InvestmentWbs             string    `xorm:"comment('投资WBS') VARCHAR(200)"`
	LeasedEquipmentIdentifier string    `xorm:"comment('租赁设备标识') VARCHAR(200)"`
	IotDeviceCode             string    `xorm:"comment('工业物联网设备编码') VARCHAR(200)"`
	CapitalizationDate        string    `xorm:"comment('转资日期') VARCHAR(200)"`
	CreatedAt                 time.Time `xorm:"DATETIME"`
	UpdatedAt                 time.Time `xorm:"DATETIME"`
	DeletedAt                 time.Time `xorm:"DATETIME"`
}
