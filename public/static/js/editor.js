var $entry_markdown_header = $("#entry-markdown-header");
var $entry_preview_header = $("#entry-preview-header");
var $entry_markdown = $(".entry-markdown");
var $entry_preview = $(".entry-preview");
var $page_name = $("#page-name");
var $page_yomi = $("#page-yomi");
var $page_message = $("#page-message");

// Tabs
$entry_markdown_header.click(function(){
  $entry_markdown.addClass('active');
  $entry_preview.removeClass('active');
});

$entry_preview_header.click(function(){
  $entry_preview.addClass('active');
  $entry_markdown.removeClass('active');
});

// $(document).on('shaMismatch', function() {
  // bootbox.dialog({
    // title: "Page has changed",
    // message: "This page has changed and differs from your draft.  What do you want to do?",
    // buttons: {
      // ignore: {
        // label: "Ignore",
        // className: "btn-default",
        // callback: function() {
          // var info = aced.info();
          // info.ignore = true;
          // aced.info(info);
        // }
      // },
      // discard: {
        // label: "Discard Draft",
        // className: "btn-danger",
        // callback: function() {
          // aced.discard();
        // }
      // },
      // changes: {
        // label: "Show Diff",
        // className: "btn-primary",
        // callback: function() {
          // bootbox.alert("Draft diff not done! Sorry");
        // }
      // }
    // }
  // });
// });

$(function(){
  $("#discard-draft-btn").click(function() {
    aced.discard();
  });

  $(".entry-markdown .floatingheader").click(function(){
    aced.editor.focus();
  });

  $("#delete-page-btn").click(function() {
    bootbox.alert("Not Done Yet! Sorry");
  });

  $('#dict-category ul li').on('click', function() {
    var innerText = $(this).text().trim();
    if (innerText === 'カテゴリー解除') {
      $('#dict-category button span').html('<span class="hidden-xs">カテゴリー <i class="fa fa-caret-down"></i></span>');
      $('#dict-category button').removeClass('btn-info').addClass('btn-default');
    } else {
      $('#dict-category button span').text(innerText);
      $('#dict-category button').removeClass('btn-default').addClass('btn-info');
    }
  });
});

var aced = new Aced({
  editor: $('#entry-markdown-content').find('.editor').attr('id'),
  renderer: function(md) { return MDR.convert(md); },
  info: Commit.info,
  submit: function(content) {
    var data = {
      name: $page_name.val(),
      yomi: $page_yomi.val(),
      message: $page_message.val(),
      content: content,
    };

    var path = Config.RELATIVE_PATH + '/' + data.name;
    var type = (Commit.info.sha) ? "PUT" : "POST";

    $.ajax({
      type: type,
      url: path,
      data: data,
      dataType: 'json'
    }).always(function(data, status, error) {
      if (status !== 'error') {
        location.href = path;
        return;
      }
      var res = data.responseJSON, r, i;

      if (data.status === 403) {
        bootbox.alert("<h3>" + "投稿が許可されていません。<br/>ログイン後再度お試し下さい。" + "</h3>");
        return;
      } else if (data.status >= 500) {
        bootbox.alert(
          "<h3>" + "投稿内容が保存できませんでした。<br/>しばらくたってから再度お試し下さい。" + "</h3>");
        return;
      }

      if (!res) {
        $page_name.addClass('parsley-error');
        bootbox.alert("<h3>" + "必須項目が入力されていません" + "</h3>");
        return;
      }

      for (i = 0; i < res.length; i++) {
        r = res[i];
        if (r.fieldNames[0] === 'yomi') $page_yomi.addClass('parsley-error');

        switch(r.message) {
        case 'Required':
          bootbox.alert("<h3>" + "必須項目が入力されていません" + "</h3>");
          break;
        default:
          bootbox.alert("<h3>" + r.message + "</h3>");
        }
      }

    });
  }
});
