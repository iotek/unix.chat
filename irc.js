var sock = new WebSocket("ws://vps.iotek.org:1337/irc");

sock.onopen = (function() {
	conslog("connected.");
	conslog("Welcome to SandBox v0.00001337a");
})

sock.onmessage = (function(msg) {
	conslog(msg.data);
})

function conslog(msg) {
	$('#console').append(msg + '<br>\n')
}

$(document).keypress(function(e) {
	if(e.which == 13) {
		sock.send(document.getElementById('prompt').value);
		document.getElementById('prompt').value = "";
	}
});
