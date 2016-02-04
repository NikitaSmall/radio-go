$(document).ready(function() {
  var client = new WebSocket("ws://localhost:3000/stream");
  client.binaryType = 'arraybuffer';

  var context = new AudioContext();
  var scriptNode = context.createScriptProcessor(16384, 1, 1);

  var bufArray = [];
  var nextStartTime = 0;
  var numSegments = 0;

  client.onopen = function() {
    console.log("established connection");
  }

  client.onmessage = function(e) {
    bufArray[bufArray.length] = e.data;
  }

  setInterval(function() {
    tempBuf = new ArrayBuffer();

    bufArray.forEach(function(item, i, arr) {
      tempBuf = arrayBufferConcat(tempBuf, item);
    });
    bufArray = [];
  }, 3000);

  setTimeout(function() {
    setInterval(function() {
      context.decodeAudioData(tempBuf, playAudio, decodeError)
    }, 3000
  )}, 3000);


  var playAudio = function(buffer) {
    source = context.createBufferSource();
    source.buffer = buffer;
    source.connect(context.destination);
    source.start();

    if (nextStartTime == 0) {
      // nextStartTime = context.currentTime;
    } else {
      // nextStartTime = nextStartTime + buffer.duration;
    }
  }

  var decodeError = function() {
    console.log("Decoding error!");
  }

  var arrayBufferConcat = function(a, b) {
    var length = 0
    var buffer = null

    length += a.byteLength + b.byteLength

    var joined = new Uint8Array(length)
    var offset = 0

    for (var i in arguments) {
      buffer = arguments[i]
      joined.set(new Uint8Array(buffer), offset)
      offset += buffer.byteLength
    }

    return joined.buffer
  }

  // Give the node a function to process audio events
  scriptNode.onaudioprocess = function(audioProcessingEvent) {
    // The input buffer is the song we loaded earlier
    var inputBuffer = audioProcessingEvent.inputBuffer;

    // The output buffer contains the samples that will be modified and played
    var outputBuffer = audioProcessingEvent.outputBuffer;

    // Loop through the output channels (in this case there is only one)
    for (var channel = 0; channel < outputBuffer.numberOfChannels; channel++) {
      var inputData = inputBuffer.getChannelData(channel);
      var outputData = outputBuffer.getChannelData(channel);

      // Loop through the 4096 samples
      for (var sample = 0; sample < inputBuffer.length; sample++) {
        // make output equal to the same as the input
        outputData[sample] = inputData[sample];

        // add noise to each output sample
        outputData[sample] += ((Math.random() * 2) - 1) * 0.2;
      }
    }
  };

});
