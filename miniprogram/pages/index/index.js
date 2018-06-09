//index.js
const app = getApp()
let Base64 = require('../../libs/js-base64/base64.js')

Page({
  data: {
    refURL: 'https://bgm.everpcpc.com',
    uid: 0,
    userInfo: null,
  },
  onLoad: function() {
    this.setData({
      'uid': app.globalData.uid,
      'userInfo': app.globalData.userInfo
    })
  },

  goSearch: app.goSearch,

  scanLogin: function(e) {
    wx.scanCode({
      success: (res) => {
        let scanData = JSON.parse(Base64.decode(res.result))

        app.globalData.chii.getUser(scanData.uid, function(res) {
          res.data.url = res.data.url.replace(/http:/, 'https:');
          Object.keys(res.data.avatar).map(function(key, index) {
            res.data.avatar[key] = res.data.avatar[key].replace(/http:/, 'https:');
          })

          app.globalData.uid = scanData.uid;
          app.globalData.token = scanData.token;
          app.globalData.userInfo = res.data;

          wx.setStorageSync('uid', scanData.uid)
          wx.setStorageSync('token', scanData.token)
          wx.setStorageSync('userInfo', res.data)

          wx.reLaunch({
            url: '../index/index',
          })
        });
      }
    })
  }
})