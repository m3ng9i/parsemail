package parsemail

import (
    "io"
    "net/mail"
)


func (email *Email) RealSender() *mail.Address {
    if email.Sender != nil {
        return email.Sender
    }

    if len(email.From) == 0 {
        return nil
    }

    return email.From[0]
}


func ParseHeader(r io.Reader) (email Email, err error) {
    msg, err := mail.ReadMessage(r)
    if err != nil {
        return
    }

    email, err = createEmailFromHeader(msg.Header)
    if err != nil {
        return
    }

    return
}


