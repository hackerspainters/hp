
function Login() {
	FB.login(function(response) {
		if (response.authResponse) {
			console.log(111)
			console.log(response.authResponse)
			console.log(response)
			console.log(222)
			// once we get the "code" and "access_token" from the authResponse, 
			// we have send them to the backend with an ajax call and let the 
			// golang backend execute the event list parsing and dump the results into
			// mongodb

			getUserInfo();
		} else {
			console.log('User cancelled login or did not fully authorize.');
		}
	},{scope: 'user_events, email'});
}  

function getUserInfo() {
	FB.api('/me', function(response) {
		var str="<b>Name</b> : "+response.name+"<br>";
		str +="<b>Link: </b>"+response.link+"<br>";
		str +="<b>Username:</b> "+response.username+"<br>";
		str +="<b>id: </b>"+response.id+"<br>";
		str +="<b>Email:</b> "+response.email+"<br>";
		str +="<input type='button' value='Get Photo' onclick='getPhoto();'/>";
		str +="<input type='button' value='Logout' onclick='Logout();'/>";
		document.getElementById("status").innerHTML=str;
	});
}

function displayUser() {
	FB.api('/me', function(response) {
		return response	
	});
}

function getPhoto() {
	FB.api('/me/picture?type=normal', function(response) {
		var str="<br/><b>Pic</b> : <img src='"+response.data.url+"'/>";
		document.getElementById("status").innerHTML+=str;
	});
}

function Logout() {
	FB.logout(function(){document.location.reload();});
}

// Load the SDK asynchronously
(function(d, s, id){
	var js, fjs = d.getElementsByTagName(s)[0];
	if (d.getElementById(id)) {return;}
	js = d.createElement(s); js.id = id;
	js.src = "//connect.facebook.net/en_US/all.js";
	fjs.parentNode.insertBefore(js, fjs);
}(document, 'script', 'facebook-jssdk'));
