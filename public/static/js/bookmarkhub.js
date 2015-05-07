$(function() {

var getTemplate, getTempateVertical;

getTemplate = function(provider, bms) {
  return "<a " + ((bms != null ? bms.link : void 0) ? "href='" + bms.link + "' target='_blank'" : void 0) + " style='margin-right: 8px;'>" +
         "  <div style='position: relative;display: inline-block;'>" +
         "    <img alt='" + provider + "' src='/static/img/icons/32/" + provider + ".png'>" +
         "    <span class='badge-mini badge' style='position: absolute; right: 0; top: 0;'>" + bms.count + "</span>" +
         "  </div>" +
         "</a>";
};

getTempateVertical = function(provider, bms) {
  return "<a " + ((bms != null ? bms.link : void 0) ? "href='" + bms.link + "' target='_blank'" : void 0) + " style='display: block; margin-bottom: 10px;'>" +
         "  <div style='position: relative;display: inline-block;'>" +
         "    <img alt='" + provider + "' src='/static/img/icons/32/" + provider + ".png' width='40'>" +
         "    <span class='badge-mini badge' style='position: absolute; right: 0; top: 0;'>" + bms.count + "</span>" +
         "  </div>" +
         "</a>";
};


$.each($('.bookmarkhub'), function() {
  var bookmarker, me;
  $(this).children().remove();
  me = this;
  bookmarker = new Bookmarkhub.Bookmarker($(this).attr('bookmarkhub-url'));
  bookmarker.twitter().done(function(data) {
    return $(me).append(getTemplate('twitter', data));
  });
  bookmarker.facebook().done(function(data) {
    return $(me).append(getTemplate('facebook', data));
  });
  bookmarker.hatena().done(function(data) {
    return $(me).append(getTemplate('hatena', data));
  });
  bookmarker.google().done(function(data) {
    return $(me).append(getTemplate('google', data));
  });
  bookmarker.pocket().done(function(data) {
    return $(me).append(getTemplate('pocket', data));
  });
  // bookmarker.delicious().done(function(data) {
    // return $(me).append(getTemplate('delicious', data));
  // });
  // bookmarker.linkedin().done(function(data) {
    // return $(me).append(getTemplate('linkedin', data));
  // });
  bookmarker.pinterest().done(function(data) {
    return $(me).append(getTemplate('pinterest', data));
  });
  // bookmarker.stumbleupon().done(function(data) {
    // return $(me).append(getTemplate('stumbleupon', data));
  // });
  // bookmarker.reddit().done(function(data) {
    // return $(me).append(getTemplate('reddit', data));
  // });
});




$.each($('.bookmarkhub-vertical'), function() {
  var bookmarker, me;
  $(this).children().remove();
  me = this;
  bookmarker = new Bookmarkhub.Bookmarker($(this).attr('bookmarkhub-url'));
  bookmarker.twitter().done(function(data) {
    return $(me).append(getTempateVertical('twitter', data));
  });
  bookmarker.facebook().done(function(data) {
    return $(me).append(getTempateVertical('facebook', data));
  });
  bookmarker.hatena().done(function(data) {
    return $(me).append(getTempateVertical('hatena', data));
  });
  bookmarker.google().done(function(data) {
    return $(me).append(getTempateVertical('google', data));
  });
  bookmarker.pocket().done(function(data) {
    return $(me).append(getTempateVertical('pocket', data));
  });
  // bookmarker.delicious().done(function(data) {
    // return $(me).append(getTempateVertical('delicious', data));
  // });
  // bookmarker.linkedin().done(function(data) {
    // return $(me).append(getTempateVertical('linkedin', data));
  // });
  bookmarker.pinterest().done(function(data) {
    return $(me).append(getTempateVertical('pinterest', data));
  });
  // bookmarker.stumbleupon().done(function(data) {
    // return $(me).append(getTempateVertical('stumbleupon', data));
  // });
  // bookmarker.reddit().done(function(data) {
    // return $(me).append(getTempateVertical('reddit', data));
  // });
});



});
