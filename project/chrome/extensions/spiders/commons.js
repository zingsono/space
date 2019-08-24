
window.MS = {};

/**
 *
 * @param seconds  休眠秒数
 * @return {Promise<unknown>}
 */
MS.sleep = async (seconds = 0)=>{
    return (await (new Promise((resolve,reject)=>{
        setTimeout(()=>{
            resolve()
        },seconds*1000);
    })));
};
