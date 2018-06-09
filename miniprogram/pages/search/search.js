// search.js
const app = getApp()

Page({
  data: {
    index: 1,
    searchTypes: [{
        id: 1,
        name: "Book"
      },
      {
        id: 2,
        name: "Anime"
      },
      {
        id: 3,
        name: "Music"
      },
      {
        id: 4,
        name: "Game"
      },
      {
        id: 6,
        name: "Real"
      }
    ]
  },
  onLoad: function(options) {},
  bindPickerChange: function(e) {
    this.setData({
      index: e.detail.value,
    })
  },
  bindSearch: function(e) {
    app.globalData.chii.search(e.detail.value.text, e.detail.value.type, function(res){
      console.log(res);
    })
  }

})