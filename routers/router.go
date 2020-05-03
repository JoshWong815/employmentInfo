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
	
	beego.Router("/getAllStudents", &controllers.StudentController{},"get:GetAllStudents")
	beego.Router("/updateStudent", &controllers.StudentController{},"*:UpdateStudent")
	beego.Router("/studentUpdating", &controllers.StudentController{},"*:StudentUpdating")
	beego.Router("/deleteStudent", &controllers.StudentController{},"*:DeleteStudent")
	beego.Router("/addStudent", &controllers.StudentController{},"*:AddStudent")
	beego.Router("/studentAdding", &controllers.StudentController{},"*:StudentAdding")
	beego.Router("/showStudents", &controllers.StudentController{},"get:ShowStudents")

	beego.Router("/getAllOffers", &controllers.OfferController{},"get:GetAllOffers")
	beego.Router("/updateOffer", &controllers.OfferController{},"*:UpdateOffer")
	beego.Router("/offerUpdating", &controllers.OfferController{},"*:OfferUpdating")
	beego.Router("/deleteOffer", &controllers.OfferController{},"*:DeleteOffer")
	beego.Router("/addOffer", &controllers.OfferController{},"*:AddOffer")
	beego.Router("/offerAdding", &controllers.OfferController{},"*:OfferAdding")

	beego.Router("/getAllEmployments", &controllers.EmploymentController{},"get:GetAllEmployments")

	beego.Router("/deleteEmployment", &controllers.EmploymentController{},"get:DeleteEmployment")

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
