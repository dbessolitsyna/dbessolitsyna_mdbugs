package tests

import (
     "github.com/revel/revel/testing"
     "net/url"
)



// Check that send works as expected and return Success
func (t *MyAppTest) TestSendButton() {
	t.Get("/")
	t.AssertOk()
	data := url.Values{}
	data.Add("message", "TestMessage")
	t.PostForm("/", data)
	t.AssertOk()
	t.AssertContains("Success")
	
}

// Check that pageTitle is "Correct page title"
func (t *MyAppTest) TestPageTitle() {
	t.Get("/")
	t.AssertOk()
	t.AssertContains("Correct page title")
	
}
