package parsemail

import "net/mail"


func (email *Email) RealSender() *mail.Address {
    if email.Sender != nil {
        return email.Sender
    }

    if len(email.From) == 0 {
        return nil
    }

    return email.From[0]
}

