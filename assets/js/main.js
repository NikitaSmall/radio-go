$(document).ready(function() {
  var player;
  var currentSong;

  $("#jquery_jplayer_1").jPlayer({
		ready: function () {
      player = this;

      $.ajax({
        method: 'GET',
        url: '/start'
      }).done(function(song) {
        currentSong = song;

        $(player).jPlayer("setMedia", {
  				title: song.title,
  				mp3: song.filePath
  			});
      });
		},

    ended: function() {
      $.ajax({
        method: 'GET',
        url: '/next/' + currentSong.id,
      }).done(function(song) {
        console.log(song);

        if (song.id) {
          currentSong = song;

          $(player).jPlayer("setMedia", {
    				title: song.title,
    				mp3: song.filePath
    			});

          $(player).jPlayer("play");
        } else {
          console.log('Playlist ended!');
        }
      });
    },
		swfPath: "/assets/js/",
		supplied: "mp3",
		wmode: "window",
		useStateClassSkin: true,
		autoBlur: false,
		smoothPlayBar: true,
		keyEnabled: true,
		remainingDuration: true,
		toggleDuration: true
	});
});
