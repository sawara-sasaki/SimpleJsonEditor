$(document).on('change', '#json-file', function() {
  var files = $(this).prop('files')
  if (files.length == 0) {
    console.log('No file selected');
    return;
  }
  const file = files[0];
  const reader = new FileReader();
  reader.onload = () => {
    SetJsonData(reader.result, false);
  };
  reader.readAsText(file);
});
var DownloadJsonFile = function() {
  var blob = new Blob(
    [$("#json-data").val()],
    { "type": "text/plain" })
  let link = document.createElement('a')
  link.href = window.URL.createObjectURL(blob)
  link.download = 'result.json'
  link.click()
}
var ReplaceNumber = function() {
  var obj = $("#json-data");
  var ori = obj.val();
  var s = obj[0].selectionStart;
  var e = obj[0].selectionEnd;
  if (e > s) {
    obj.val(ori.substring(0, s) + replaceNumber(ori.substring(s, e), $("#replace-number").val()) + ori.substring(e));
  }
}
var replaceNumber = function(str, num) {
  return str.replace(/\d+/g, num);
}
var AddLinearData = function() {
  const data = {action:"linear", parameters:[$("#linear-start").val(), $("#linear-end").val()]}
  request(data, , (res)=>{
    AddData(res.data);
  }, onerror);
}
var AddData = function(data) {
  var obj = $("#json-data");
  var ori = obj.val();
  var s = obj[0].selectionStart;
  obj.val(ori.substring(0, s) + data + ori.substring(s));
}
var onError = function(e) {
  if (!!e.responseJSON) {
    console.log(e.responseJSON.message);
    showDangerMessage(e.responseJSON.message);
  } else {
    console.log(e.message);
    showDangerMessage(e.message);
  }
};
var request = function(data, callback, onerror) {
  $.ajax({
    type:          'POST',
    dataType:      'json',
    contentType:   'application/json',
    scriptCharset: 'utf-8',
    data:          JSON.stringify(data),
    url:           App.url
  })
  .done(function(res) {
    callback(res);
  })
  .fail(function(e) {
    onerror(e);
  });
};
var CheckSyntax = function() {
  var data = $("#json-data").val();
  if (!data || data.length < 1) {
    showDangerMessage("No Data.");
  } else {
    SetJsonData(data, true);
  }
}
var SetJsonData = function(data, showInfo) {
  try {
    const json = JSON.parse(data);
    if (!Object.keys(json).length) {
      showDangerMessage("No Data.");
    } else {
      if (showInfo) {
        showInfoMessage("Syntax OK.");
      }
      $("#json-data").val(JSON.stringify(json, null, "\t"));
    }
  } catch (error) {
    showDangerMessage("Syntax Error.");
  }
}
var showDangerMessage = function(str) {
  $("#message-danger").text(str);
  $("#message-danger-container").removeClass("hide");
}
var showInfoMessage = function(str) {
  $("#message-info").text(str);
  $("#message-info-container").removeClass("hide");
}
var CloseMessage = function(elm) {
  $(elm).parent().addClass("hide");
}
