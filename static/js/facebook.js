
function Login() {
	FB.login(function(response) {
		if (response.authResponse) {
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

function RegisterAttendee(eid) {
	FB.login(function(response) {
		if (response.authResponse) {
			FB.api('/me', function(resp) {
				data = JSON.stringify({
					eid: eid, fbuid: resp.id, firstname: resp.first_name, lastname: resp.last_name, email: resp.email
				})
				$.ajax({
					type: 'POST',
					url: "/events/register/",
					data: data,
					dataType: 'json',
					contentType: 'application/json'
				});
			});
		} else {
			console.log('Event registration failed.');
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
		var status = document.getElementById("status"); 
		if (status != null) {
			document.getElementById("status").innerHTML=str;
		}
	});
}

function displayUser() {
	FB.api('/me', function(response) {
		$("#express").text("Attending as "+ response.first_name)
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

function getGroupEvents(fbuid, gid, token, expiresIn) {

	FB.api('/'+gid+'?access_token='+token, function(response) {
		if (response.owner.id == fbuid) {

			data = JSON.stringify({token: token, expiresIn: expiresIn}),

			$.ajax({
				type: 'POST',
				url: "/events/import/",
				data: data,
				dataType: 'json',
				contentType: 'application/json'
			});

			FB.api('/'+gid+'/events?access_token='+token, function(response) {
				for (var i=0; i<response.data.length; i++) {
					getEvent(response.data[i].id, token)
				}
			});
		}
	});

}

function getEvent(eid, token) {

	FB.api('/'+eid+'?access_token='+token, function(response) {
		// no callback needed
	});

}

// Load the SDK asynchronously
(function(d, s, id){
	var js, fjs = d.getElementsByTagName(s)[0];
	if (d.getElementById(id)) {return;}
	js = d.createElement(s); js.id = id;
	js.src = "//connect.facebook.net/en_US/all.js";
	fjs.parentNode.insertBefore(js, fjs);
}(document, 'script', 'facebook-jssdk'));
