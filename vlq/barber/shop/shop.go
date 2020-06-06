// Package shop implements the sleeping barber problem.
// There is one barber in the barber shop, one barber chair and n chairs for
// waiting customers. If there are no customers, the barber sits down in the
// barber chair and takes a nap. An arriving customer must wake the barber.
// Subsequent arriving customers take a waiting chair if any are empty or
// leave if all chairs are full.
//
// Have the ability to close the shop even if new customers are entering.
// Customers looking for a chair should run on their own goroutine.
//
// Task: Change EnterCustomer so a customer can wait for a specified amount
// of time for a chair to open up.
