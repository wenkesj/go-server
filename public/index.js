(function(document, window, $) {
  $(document).ready(function() {
    var textArea = document.getElementById('in-js');
    textArea.value = "// Write some javascript here and press compile!\n";
    var editor = new MirrorFrame(CodeMirror.replace(textArea), {
      height: "350px",
      content: textArea.value,
      parserfile: ["/javascripttokenizer.js", "/javascriptparser.js"],
      stylesheet: "/javascriptstyle.css",
      autoMatchParens: true
    });
    var form = $('#compile-js');
    form.submit(function(e) {
      e.preventDefault();
      var javascript = {
        javascript: editor.mirror.editor.getCode()
      };
      $.ajax('/js', {
        data : JSON.stringify(javascript),
        contentType : 'application/json',
        type : 'POST',
        success: function(compiledJs) {
          $('#out-js').text(compiledJs);
          return;
        }
      });
      return false;
    });
  });
})(document, window, $);
