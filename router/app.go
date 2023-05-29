package router

import (
	"github.com/gin-gonic/gin"
	"hie/main/service"
)

func Router() (engine *gin.Engine) {
	r := gin.Default()
	//路由
	r.POST("login", service.Login)
	r.POST("create", service.Create)
	//用户组
	user := r.Group("user")
	{
		user.POST("create", service.UserDataReceive) //添加用户
		user.DELETE("delete", service.UserDelete)    //删除用户
		user.GET("getallusers", service.GetAllUser)  //获取全部用户数据
		user.POST("detailchange", service.UserDetailChange)
		user.POST("select", service.UserSelect) //模糊查询
	}

	//医生组
	doctor := r.Group("doctor")
	{
		doctor.POST("create", service.DoctorDataReceive) //添加医生
		doctor.DELETE("delete", service.DoctorDelete)    //删除医生
		doctor.GET("getall", service.GetAllDoctor)
		doctor.POST("detailchange", service.DoctorDetailChange)
		doctor.POST("select", service.DoctorSelect)                     //
		doctor.POST("getDoctorByDepartId", service.GetDoctorByDepartId) //通过科室id查医生
	}

	//科室组
	department := r.Group("department")
	{
		department.POST("create", service.DepartmentCreate)
		department.DELETE("delete", service.DepartmentDelete)
		department.GET("getall", service.DepartmentGet)
		department.POST("detailchange", service.DepartmentDetail)
		department.POST("select", service.DepartmentSelect) //
	}

	//会员组
	member := r.Group("member")
	{
		member.POST("create", service.MemberDataReceive)        //添加会员
		member.DELETE("delete", service.MemberDelete)           //删除会员
		member.GET("get", service.MemberGet)                    //获取所有
		member.POST("detailchange", service.MemberDetailChange) //修改信息
		member.POST("select", service.MemberSelect)             //模糊查询
		member.POST("recharge", service.MemberRecharge)         //会员充值
	}

	//药品组
	drug := r.Group("drug")
	{
		drug.POST("create", service.DrugCreate)
		drug.DELETE("delete", service.DrugDelete)
		drug.GET("get", service.DrugGet)
		drug.POST("detailchange", service.DrugDetailChange)
		drug.POST("select", service.DrugSelect)
		drug.POST("increasequantity", service.IncreaseQuantity) //药品数量增加
	}

	//挂号组
	registerorder := r.Group("registerorder")
	{
		registerorder.POST("create", service.RegisterOrderCreate)
		registerorder.DELETE("delete", service.RegisterOrderDelete)
		registerorder.GET("get", service.RegisterOrderGet)
		registerorder.POST("select", service.RegisterOrderSelect)
	}

	//医生值班组
	doctorduty := r.Group("doctorduty")
	{
		doctorduty.POST("create", service.DoctordutyCreate)
		doctorduty.DELETE("delete", service.DoctordutyDelete)
		doctorduty.GET("get", service.DoctordutyGet)
		doctorduty.POST("select", service.DoctordutySelect)
	}

	//医嘱
	doctoradvice := r.Group("doctoradvice")
	{
		doctoradvice.POST("create", service.DoctorAdviceCreate)
		doctoradvice.DELETE("delete", service.DoctorAdviceDelete)
		doctoradvice.GET("get", service.DoctorAdviceGet)
		doctoradvice.POST("select", service.DoctorAdviceSelect)
		doctoradvice.GET("getall", service.DoctorAdviceByNotGrantGetAll)
	}
	//开药
	prescribe := r.Group("prescribe")
	{
		prescribe.POST("create", service.PrescribeCreate)
		prescribe.DELETE("delete", service.PrescribeDelete)
		prescribe.POST("getbydaid", service.PrescribeGetBydaid)
		prescribe.POST("grant", service.PrescribeGrant)
	}

	//缴费明细
	paymentDetails := r.Group("paymentDetails")
	{
		paymentDetails.GET("getall", service.PaymentGet)
		paymentDetails.POST("select", service.PaymentDetailsSelect)
	}

	//费用结算
	costsettledetails := r.Group("costsettledetails")
	{
		costsettledetails.GET("getallNotSettlement", service.GetallNotSettlement)
	}
	return r
}
