(function(global) {

  'use strict';

  var global = global || {};

  // existing version for noConflict()
  var _Chii = global.Chii;

  var version = "1.0.0";

  var apiURL = "https://api.bgm.tv";
  var appID = "bgm225a93867f219bb";
  var token = "";
  var uid = 0;

  var init = function(_uid, _token) {
    uid = _uid;
    token = _token;
  }

  var getUser = function(uid, success) {
    wx.request({
      url: apiURL + '/user/' + uid,
      success: success
    })
  }

  var search = function(keywords, type, success) {
    wx.request({
      url: apiURL + '/search/subject/' + encodeURI(keywords),
      data: {
        type: type,
      },
      success: success
    })
  }

  var getSubject = function(id, success) {
    wx.request({
      url: apiURL + '/subject/' + id,
      success: success
    })
  }

  var userCollections = function(subjectType, success) {
    wx.request({
      url: apiURL + '/user/' + uid + '/collections/'+ subjectType,
      data: {
        app_id: appID,
      },
      success: success,
    })
  }

  var userProgress = function (subjectID, success) {
    wx.request({
      url: apiURL + '/user/' + uid + '/progress',
      header: {
        Authorization: "Bearer "+ token,
      },
      data: {
        app_id: appID
      },
      success: success,
    })
  }

  // export Chii
  global.Chii = {
    VERSION: version,
    init: init,
    getUser: getUser,
    getSubject: getSubject,
    userProgress: userProgress,
    userCollections: userCollections,
    search: search
  };
  module.exports = global.Chii;
})(this);