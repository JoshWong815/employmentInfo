package routers

import (
	"employmentInfo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"*:Login")
	beego.Router("/login", &controllers.MainController{},"*:Login")
	beego.Router("/loginTest", &controllers.MainController{},"*:LoginTest")
	beego.Router("/logout", &controllers.MainController{},"*:Logout")
	beego.Router("/index", &controllers.MainController{},"*:Index")


	beego.Router("/studentMainPage", &controllers.SingleStudentController{},"*:StudentMainPage")
	beego.Router("/singleStudentInfo", &controllers.SingleStudentController{},"*:SingleStudentInfo")
	beego.Router("/singleStudentUpdating", &controllers.SingleStudentController{},"*:SingleStudentUpdating")



	beego.Router("/getAllAdmins", &controllers.AdminController{},"get:GetAllAdmins")
	beego.Router("/updateAdmin", &controllers.AdminController{},"*:UpdateAdmin")
	beego.Router("/AdminUpdating", &controllers.AdminController{},"*:AdminUpdating")
	beego.Router("/deleteAdmin", &controllers.AdminController{},"*:DeleteAdmin")
	beego.Router("/addAdmin", &controllers.AdminController{},"*:AddAdmin")
	beego.Router("/adminAdding", &controllers.AdminController{},"*:AdminAdding")
	beego.Router("/showAdmins", &controllers.AdminController{},"get:ShowAdmins")
	
	beego.Router("/getAllStudents", &controllers.StudentController{},"get:GetAllStudents")
	beego.Router("/updateStudent", &controllers.StudentController{},"*:UpdateStudent")
	beego.Router("/studentUpdating", &controllers.StudentController{},"*:StudentUpdating")
	beego.Router("/deleteStudent", &controllers.StudentController{},"*:DeleteStudent")
	beego.Router("/addStudent", &controllers.StudentController{},"*:AddStudent")
	beego.Router("/studentAdding", &controllers.StudentController{},"*:StudentAdding")
	beego.Router("/showStudents", &controllers.StudentController{},"get:ShowStudents")
	beego.Router("/fileUpload", &controllers.StudentController{},"*:FileUpload")


	beego.Router("/getAllOffers", &controllers.OfferController{},"get:GetAllOffers")
	beego.Router("/updateOffer", &controllers.OfferController{},"*:UpdateOffer")
	beego.Router("/offerUpdating", &controllers.OfferController{},"*:OfferUpdating")
	beego.Router("/deleteOffer", &controllers.OfferController{},"*:DeleteOffer")
	beego.Router("/addOffer", &controllers.OfferController{},"*:AddOffer")
	beego.Router("/offerAdding", &controllers.OfferController{},"*:OfferAdding")

	beego.Router("/getAllEmployments", &controllers.EmploymentController{},"get:GetAllEmployments")
	//beego.Router("/updateEmployments", &controllers.EmploymentController{},"*:UpdateEmployment")
	beego.Router("/deleteEmployment", &controllers.EmploymentController{},"get:DeleteEmployment")
	beego.Router("/addEmployment", &controllers.EmploymentController{},"*:AddEmployment")
	beego.Router("/employmentAdding", &controllers.EmploymentController{},"*:EmploymentAdding")
	beego.Router("/getNowCompany", &controllers.EmploymentController{},"*:GetNowCompany")
	beego.Router("/checkSid", &controllers.EmploymentController{},"*:CheckSid")
	beego.Router("/getSidEmployment", &controllers.EmploymentController{},"*:GetSidEmployment")
	beego.Router("/getAllCompanyNames", &controllers.EmploymentController{},"*:GetAllCompanyNames")


	beego.Router("/getAllCompanys", &controllers.CompanyController{},"get:GetAllCompanys")
	beego.Router("/deleteCompany", &controllers.CompanyController{},"*:DeleteCompany")
	beego.Router("/updateCompany", &controllers.CompanyController{},"*:UpdateCompany")
	beego.Router("/companyUpdating", &controllers.CompanyController{},"*:CompanyUpdating")
	beego.Router("/addCompany", &controllers.CompanyController{},"*:AddCompany")
	beego.Router("/companyAdding", &controllers.CompanyController{},"*:CompanyAdding")

	beego.Router("/getAllSkills", &controllers.SkillController{},"get:GetAllSkills")
	beego.Router("/deleteSkill", &controllers.SkillController{},"*:DeleteSkill")
	beego.Router("/updateSkill", &controllers.SkillController{},"*:UpdateSkill")
	beego.Router("/skillUpdating", &controllers.SkillController{},"*:SkillUpdating")
	beego.Router("/addSkill", &controllers.SkillController{},"*:AddSkill")
	beego.Router("/skillAdding", &controllers.SkillController{},"*:SkillAdding")
}
