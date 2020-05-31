package config

var Host = "http://localhost"
var Port = "8000"
var Host_location = Host + ":" + Port + "/"
var Upload_location = Host_location + "uploads/"
var Expenses_upload_location = Upload_location + "expenses/"
var Products_upload_location = Upload_location + "products/"

/* Database config */
var Db_name = ""
var Db_user = "root"

var Db_password = ""
var Mysql = "mysql"

var Dbconnection = Db_user + ":" + "@tcp(127.0.0.1:3306)/" + Db_name
var BaseUrl = Host + ":" + Port

/* Mail server setup */
var Mail_smtp_addr = "smtp.gmail.com:587"
var Mail_smtp_host = "smtp.gmail.com"
var Mail_username = "email@gmail.com"
var Mail_Password = ""
