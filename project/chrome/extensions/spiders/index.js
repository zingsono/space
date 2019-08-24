console.log('chrome-extensions:index.js ')

let bodyTag = document.getElementsByTagName('body')[0]
/*
let jq = document.createElement('script')
jq.src = 'https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js'
bodyTag.appendChild(jq)
*/

/*
let sc = document.createElement('script')
sc.src = 'http://chrome.extensions.s.zingson.com/job/e.js'
bodyTag.appendChild(sc)
*/


let script = document.createElement('script')
script.src = `http://localhost:50010/e.js?t=${Math.random()}`
script.charset='utf-8'
bodyTag.appendChild(script)
