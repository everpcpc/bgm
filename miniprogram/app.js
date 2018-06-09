//app.js

let Chii = require('libs/chobits/chii.js')

App({
  globalData: {
    uid: wx.getStorageSync('uid') || 0,
    token: wx.getStorageSync('token') || null,
    userInfo: wx.getStorageSync('userInfo') || null,
    chii: null
  },
  onLaunch: function() {
    if (this.globalData.uid) {
      this.globalData.chii = Chii
      this.globalData.chii.init(this.globalData.uid, this.globalData.token.access_token)
    }
  },

  goSearch: function (e) {
    wx.navigateTo({
      url: '../search/search',
    })
  }

})