package udpLog

//应用程序
type BootApplication struct {
	IsTest     bool
}

//构造系统
func New() *BootApplication {
	e := &BootApplication{}
	return e
}

func (e *BootApplication) Start() {
	//1. 初始化starter
	e.init()
	//2. 安装starter
	e.setup()
	//3. 启动starter
	e.start()
}

//程序初始化
func (e *BootApplication) init() {
	for _, v := range GetStarters() {
		v.Init()
	}
}

//程序安装
func (e *BootApplication) setup() {

	//log.Info("Setup starters...")
	for _, v := range GetStarters() {
		//typ := reflect.TypeOf(v)
		//log.Debug("Setup: ", typ.String())
		v.Setup()
	}

}

//程序开始运行，开始接受调用
func (e *BootApplication) start() {

	//log.Info("Starting starters...")
	for i, v := range GetStarters() {

		//typ := reflect.TypeOf(v)
		//log.Debug("Starting: ", typ.String())
		if v.StartBlocking() {
			if i+1 == len(GetStarters()) {
				v.Start()
			} else {
				go v.Start()
			}
		} else {
			v.Start()
		}

	}
}

//程序开始运行，开始接受调用
func (e *BootApplication) Stop() {

	//log.Info("Stoping starters...")
	for _, v := range GetStarters() {
		//typ := reflect.TypeOf(v)
		//log.Debug("Stoping: ", typ.String())
		v.Stop()
	}
}
