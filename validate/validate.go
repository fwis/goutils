package validate

import
//"fmt"

"regexp"

// 只能是 汉字、英文字母、数字、下划线
func IsNormalStr(s string) bool {
	re := regexp.MustCompile(`^[\x{0391}-\x{FFE5}\w\-\.,]+$`)
	return re.MatchString(s)
}

// 只能是英文字母、数字、下划线、减号、点号
func IsNormalChars(s string) bool {
	re := regexp.MustCompile(`^[\w\-\.]+$`)
	return re.MatchString(s)
}

func IsZHMobile(s string) bool {
	re := regexp.MustCompile(`^1[34578]{1}\d{9}$`)
	return re.MatchString(s)
}

func IsZHPhone(s string) bool {
	re := regexp.MustCompile(`^((\+86)|(86))?(\-)?(0[0-9]{2,3}\-)?([2-9][0-9]{6,7})+(\-[0-9]{1,6})?$`)
	return re.MatchString(s)
}

func IsEmail(s string) bool {
	re := regexp.MustCompile("^[_a-zA-Z0-9-]+(\\.[_a-zA-Z0-9-]+)*@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*(\\.[a-zA-Z]{2,3})$")
	return re.MatchString(s)
}
