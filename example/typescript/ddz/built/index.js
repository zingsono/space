"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const _ = require("lodash");
//0-53 代表54张牌
//let poker:Array = [{label:'大王',bg:'/0.png',audio:'0.mp3'}]
// 54张牌 编号0-53分别对应为：红3、方3、黑3、花3、红4......、红A、方A、黑A、花A、红2、方2、黑2、花2、小王、大王
let cards = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53];
// 54张牌根据大小分为14组，最后一组[52,53]两张牌,代表小王、大王
let cardGroup = _.chunk(cards, 4);
//已经出牌的集合，记录已经出的牌
let discard = [];
//当前出牌集合,记录当前用户出牌，下个用户出牌必须比当前用户大
let currentCard = [];
//发牌：返回4个数组，前三个数组17张代表3个人的手牌，第四个数组3张代表底牌
let drawCard = (cards) => {
    //拷贝一份数组
    let randomCards = Object.assign({}, cards);
    //打乱数组顺序
    randomCards = _.shuffle(randomCards);
    console.log('随机发牌集合：', JSON.stringify(randomCards));
    //分为3组，每组取前17张做手牌，最后一张作为底牌
    let d = _.chunk(randomCards, 18);
    console.log('_.chunk = ', JSON.stringify(d));
    let dp = [_.takeRight(d[0], 1)[0], _.takeRight(d[1], 1)[0], _.takeRight(d[2], 1)[0]];
    let s = [];
    s.push(_.take(d[0], 17));
    s.push(_.take(d[1], 17));
    s.push(_.take(d[2], 17));
    s.push(dp);
    return s;
};
let drawCardResult = drawCard(cards);
console.log('发牌结果：', JSON.stringify(drawCardResult));
let user = {
    'U001': drawCardResult[0].sort((a, b) => b - a),
    'U002': drawCardResult[1].sort((a, b) => b - a),
    'U003': drawCardResult[2].sort((a, b) => b - a),
};
let owner = 'U001'; //地主用户
user[owner].push(...drawCardResult[3]);
user[owner] = user[owner].sort((a, b) => b - a);
console.log('用户手牌：', JSON.stringify(user));
//出单张牌判断
let C1 = (uid, ...ns) => {
};
