package parsemail

import (
    "fmt"
    "io"
    "mime"
    "net/mail"

    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
)


var wordDecoder = mime.WordDecoder {
    CharsetReader : func(charset string, input io.Reader) (reader io.Reader, err error) {
        switch charset {
            case "gb2312":
                // use simplifiedchinese.HZGB2312 cannot get the expected result
                reader = transform.NewReader(input, simplifiedchinese.GBK.NewDecoder())
            case "gbk":
                reader = transform.NewReader(input, simplifiedchinese.GBK.NewDecoder())
            case "gb18030":
                reader = transform.NewReader(input, simplifiedchinese.GB18030.NewDecoder())
            default:
                err = fmt.Errorf("unknown charset: '%s'", charset)
                return
        }
        return
    },
}


// AddressParser is for parse address with name contains chinese character.
var AddressParser = mail.AddressParser { WordDecoder : &wordDecoder }

