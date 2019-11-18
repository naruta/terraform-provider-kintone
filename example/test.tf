provider "kintone" {
}

resource "kintone_application" "test_application_1" {
  name        = "Terraform MyApp"
  description = "This is my application!!"
  theme       = "BLACK"

  field {
    code  = "code_number"
    label = "number"
    type  = "NUMBER"
  }
  field {
    code  = "code_title"
    label = "title"
    type  = "SINGLE_LINE_TEXT"
  }
  field {
    code  = "code_description"
    label = "description"
    type  = "MULTI_LINE_TEXT"
  }

  status_enable = true

  state {
    name  = "New"
    index = 0
  }
  state {
    name  = "InProgress"
    index = 1
  }
  state {
    name  = "Done"
    index = 2
  }

  action {
    name = "InProgress!!"
    from = "New"
    to   = "InProgress"
  }
  action {
    name = "Done!!Done"
    from = "InProgress"
    to   = "Done"
  }
}

# resource "kintone_record" "test_record_1" {
#   app_id = kintone_application.test_application_1.id
#
#   values = <<VALUES
# {
#     "code_number": "12345",
#     "code_title": "test title"
# }
# VALUES
#
# }

# resource "kintone_application" "test_application_2" {
#   name        = "Terraform MyApp 2"
#   description = "This is my application!!"
#   theme       = "BLUE"
#
#   field {
#     code  = "code_app"
#     label = "AppId"
#     type  = "SINGLE_LINE_TEXT"
#   }
#
#   view {
#     name   = "Test View"
#     index  = 0
#     type   = "LIST"
#     fields = ["code_app"]
#   }
# }

# resource "kintone_record" "test_record_2" {
#   app_id = kintone_application.test_application_2.id
#
#   values = <<VALUES
# {
#     "code_app": "${kintone_application.test_application_1.id}"
# }
# VALUES
#
# }
#
