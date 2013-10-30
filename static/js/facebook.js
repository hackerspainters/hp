window.fbAsyncInit = function() {
	// init the FB JS SDK
	FB.init({
		//appId      : '521976014544775',                    // App ID from the app dashboard
		//channelUrl : '//hackersandpainters.sg/channel.html', // Channel file for x-domain comms

		// for local machine testing purposes
		 appId      : '232276926931495',
		 channelUrl : '//calvinchengx.dyndns.org/channel.html', 

		status     : true,                                 // Check Facebook Login status
		xfbml      : true                                  // Look for social plugins on the page
	});

	// Additional initialization code such as adding Event Listeners goes here

	FB.Event.subscribe('auth.authResponseChange', function(response) {
		if (response.status === 'connected') {
			console.log("connected")
			console.log(response)
			console.log(response.authResponse.accessToken)
			document.getElementById("message").innerHTML +=  "<br>Connected to Facebook";
			result = displayUser()
			console.log(result)
			$("#status").html("Connected " + result.name)
			//SUCCESS

		} else if (response.status === 'not_authorized') {
			document.getElementById("message").innerHTML +=  "<br>Failed to Connect";
			$("#status").html("Not connected (1)")
			//FAILED
		} else {
			document.getElementById("message").innerHTML +=  "<br>Logged Out";
			$("#status").html("Not connected (2)")
			//UNKNOWN ERROR
		}
	});	

};

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
