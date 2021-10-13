// Copyright (C) 2015  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package hrpc

// Notifier represents something that will be notified
// before and after a Call is actually sent
type Notifier interface {
	BeforeSend()
	AfterReceive()
	Done(err error)
}

// NotificationSend represents objects that can notify
// Notifiers when they have been sent
type NotificationSender interface {
	BeforeSend()
	AfterReceive()
	Done(err error)
	AddNotifier(n Notifier)
}
