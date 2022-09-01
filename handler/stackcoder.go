package handler

import (
	"context"
 
	log "github.com/micro/micro/v3/service/logger"

	stackcoder "stackcoder/proto"
)

type Stackcoder struct{}

// Return a new handler
func New() *Stackcoder {
	return &Stackcoder{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Stackcoder) Call(ctx context.Context, req *stackcoder.Request, rsp *stackcoder.Response) error {
	log.Info("Received Stackcoder.Call request")
	
	rsp.Msg = "Hello " + req.Name
	return nil
}
func (e *Stackcoder) Hints(ctx context.Context, req *stackcoder.HintsRequest, rsp *stackcoder.Response) error {
	log.Info("Received Stackcoder.Call request")
	rsp.Msg = `{
		"Hints": [
			{"Title": "How to get the value of checked mat-radio-button in Angular? " , "Links": "https://stackoverflow.com/questions/73564803/how-to-get-the-value-of-checked-mat-radio-button-in-angular" , "Notes": "I am building an angular page with a table having a radio button","Hints": "0 Answers Yet"},
			{"Title": "Angular 14: Subtract 1 day when converting data toISOString() " , "Links": "https://stackoverflow.com/questions/73562967/angular-14-subtract-1-day-when-converting-data-toisostring" , "Notes": "I have a calendar in a parent component, and I pass the selected date as a parameter via selector to a child component:","Hints": "Without converting date to ISO format, it will display your local time (with time zone included).

			ISO format is independent of the time zone.
			
			You should consider doing everything in ISO format, especially if you need to send data to the server for processing."}, 
		]

	}`
	return nil
}
func (e *Stackcoder) Users(ctx context.Context, req *stackcoder.UsersRequest, rsp *stackcoder.Response) error {
	log.Info("Received Stackcoder.Call request") 
	rsp.Msg = `{
		"Users": [
			{"Title": "Use " , "email": "admin@mail.com" , "password": "password"},
			{"name": "User" , "email": "user@mail.com" , "password": "password"},
		]

	}`
	return nil
}
 
func (e *Stackcoder) Tutorials(ctx context.Context, req *stackcoder.TutorialsRequest, rsp *stackcoder.Response) error {
	log.Info("Received Stackcoder.Call request")
	rsp.Msg = `{
		"Tutorials": [
			{"title": "Go Oauth2 Tutorial" , "content": "Welcome fellow coders! In this tutorial, we are going to be taking a look at how you can implement your own OAuth2 Server and client using the go-oauth2/oauth2 package.

This is without a doubt one of the most requested topics from commentors on my YouTube videos and its certainly something that I myself find incredibly interesting.

Security is without doubt a very important feature for any public and even private facing service or API and its something that you need to pay a lot of attention to in order to get it right.

Note - The full github repository for this tutorial can be found here: TutorialEdge/go-oauth-tutorial

The Theory
So, before we dive into how we can code this up, its important to know how it works in the background. Typically, we have a client that will start by making an authorization request to the resource owner. The resource owner then either grants or denies this request.

With this authorization grant the client then passes this to the authorization server which will grant back an access token. It is with this granted access token that our client can then access a protected resource such as an API or a service.

So, with that said, lets now look at how we can implement our own authorization server using this go-oauth2/oauth2 package." , "author": "Tutorialedge" ,"link": "https://tutorialedge.net/golang/go-oauth2-tutorial/", "image": "" },
{"title": "Go Oauth2 Tutorial" , "content": "Welcome fellow coders! In this tutorial, we are going to be taking a look at how you can implement your own OAuth2 Server and client using the go-oauth2/oauth2 package.

This is without a doubt one of the most requested topics from commentors on my YouTube videos and its certainly something that I myself find incredibly interesting.

Security is without doubt a very important feature for any public and even private facing service or API and its something that you need to pay a lot of attention to in order to get it right.

Note - The full github repository for this tutorial can be found here: TutorialEdge/go-oauth-tutorial

The Theory
So, before we dive into how we can code this up, its important to know how it works in the background. Typically, we have a client that will start by making an authorization request to the resource owner. The resource owner then either grants or denies this request.

With this authorization grant the client then passes this to the authorization server which will grant back an access token. It is with this granted access token that our client can then access a protected resource such as an API or a service.

So, with that said, lets now look at how we can implement our own authorization server using this go-oauth2/oauth2 package." , "author": "Tutorialedge" ,"link": "https://tutorialedge.net/golang/go-oauth2-tutorial/", "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQVvXRW2pfa_IJ2LoVM7c9YDhmq8tn7xWxb3ljPIDF8&s" },
{"title": "Building a simple microservices using Go Micro and Echo frameworks" , "content": "Introduction
In this article, we will learn how to build a simple microservices using Go Micro and Echo frameworks. And with Go Micro, building a microservices has never been easy than this before.

Building our first microservices
Generate services
We will generate two services here: the first one is API gateway and the other one is a simple service to demonstrate calling between our services." , "author": "Cozi.Dev" ,"link": "https://cozi.dev/posts/building-a-simple-microservices-using-go-micro-and-echo-frameworks/", "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQVvXRW2pfa_IJ2LoVM7c9YDhmq8tn7xWxb3ljPIDF8&s" },
{"title": "Go Angular Tutorial" , "content": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum." , "author": "Cozi.Dev" ,"link": "https://cozi.dev/posts/building-a-simple-microservices-using-go-micro-and-echo-frameworks/", "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQVvXRW2pfa_IJ2LoVM7c9YDhmq8tn7xWxb3ljPIDF8&s" },
		]

	}`
	return nil
}