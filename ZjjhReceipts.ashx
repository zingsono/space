<%@ WebHandler Language="C#" Class="UnionLive.ShopService.Activity.ZjjhReceipts" %>

using System;
using System.Collections.Generic;
using System.Drawing;
using System.IO;
using System.Runtime.InteropServices;
using System.Text;
using System.Web;
using System.Xml;
using UnionLive.Common;
using UnionLive.Framework;
using UnionLive.SmsService.Core;
using TransactionException = UnionLive.Basic.TransactionException;


namespace UnionLive.ShopService.Activity
{
    /// <summary>
    /// 浙江建行代收
    /// </summary>
    public class ZjjhReceipts : IHttpHandler
    {
        private const string DllAddr = @"U:\UnionLive\libs\bhs\bhs.dll";
        private string host = UnionLive.Shop.SysParams.Default.RequstHostId;
        private int port = UnionLive.Shop.SysParams.Default.RequestPortId;
        private string projectNo = UnionLive.Shop.SysParams.Default.JhProjectNo;
        private string companyNo = UnionLive.Shop.SysParams.Default.JhCompanyNo;
        public void ProcessRequest(HttpContext context)
        {
            var request = context.Request;
            Logger.LogInfo(string.Format("收到交易报文{0}", request.RawUrl));
            var servcieType = request["t"];
            string result = "";
            try
            {
                switch (servcieType)
                {
                    case "test":
                        result = Test(request["name"], request["cardNo"], request["idNo"]
                            , request["mobile"]);
                        break;
                    //4.7 可签约扣款地区查询-WXDK07
                    case "QueryDeductMoneyArea":
                        result = QueryDeductMoneyArea();
                        break;
                    //4.8 代扣项目查询-WXDK08
                    case "QueryDeductMoneyProject":
                        result = QueryDeductMoneyProject(request["areaId"]);
                        break;
                    //4.1 代扣签约申请-WXDK01
                    case "DeductMoneyProjectApply":
                        result = DeductMoneyProjectApply(request["entno"], request["idx"], request["idxname"]
                            , request["actno"], request["id"], request["mobile"], request["waterMark"]);
                        break;
                    //4.2 代扣签约验证生效-WXDK02
                    case "ProjectApplyValid":
                        result = ProjectApplyValid(request["entno"], request["idx"], request["actno"]
                            , request["mobile"], request["validcode"]);
                        break;
                    //4.3 代扣签约撤销申请-WXDK03
                    case "ProjectCancelApply":
                        result = ProjectCancelApply(request["entno"], request["idx"], request["actno"]
                            , request["mobile"], request["waterMark"]);
                        break;
                    //4.4 代扣签约撤销验证生效-WXDK04
                    case "ProjectCancelValid":
                        result = ProjectCancelValid(request["entno"], request["idx"], request["actno"]
                            , request["mobile"], request["validcode"]);
                        break;
                    //4.5 代扣签约查询申请-WXDK05
                    case "ProjectSearchApply":
                        result = ProjectSearchApply(request["idx"], request["idxname"]
                            , request["actno"], request["id"], request["mobile"], request["waterMark"]);
                        break;
                    //4.6 代扣签约查询验证返回-WXDK06
                    case "ProjectSearchValid":
                        result = ProjectSearchValid(request["idx"], request["actno"]
                            , request["mobile"], request["validcode"]);
                        break;
                    case "GetValidCode":
                        result = GetValidCode(request["mobile"]);
                        break;
                    //发送短信
                    case "SendSms":
                        result = SendSms(request["phone"], request["smsTemplateId"]);
                        break;
                }
            }
            catch (TransactionException ex)
            {
                Logger.LogException(ex);
                result = "{\"returnCode\":\"" + ex.ErrorCode + "\",\"returnMessage\":\"" + ex.Message + "\"}";
            }
            catch (Exception ex)
            {
                Logger.LogException(ex);
                result = "{\"returnCode\":\"9999\",\"returnMessage\":\"系统繁忙,请稍后再试\"}";
            }
            Logger.LogInfo(string.Format("返回报文:{0}", result));
            context.Response.Write(result);
        }


        /// <summary>
        /// 发送短信
        /// </summary>
        public string SendSms(string phone, string templateId = null)
        {
            var random = new Random();
            var code = random.Next(100000, 999999).ToString();
            var smsManager = new SmsManager();
            var codeSeq = random.Next(10, 99).ToString();
            var smsContent = "验证码:" + code + "，序号:" + codeSeq;
            var hostTrace = Basic.Retrieval.GetNextTraceNo();
            if (Checker.IsNotEmpty(templateId))
            {
                var rv = new Shop.Retrieval();
                var ds = rv.RvSmsPlate(templateId);
                if (Checker.IsNotEmpty(ds))
                {
                    var nowSmsContent = ds.PickValue("TEMPLATE_CONTENT");
                    if (Checker.IsNotEmpty(nowSmsContent))
                    {
                        smsContent = nowSmsContent;
                        smsContent = smsContent.Replace("{验证码}", code);
                        smsContent = smsContent.Replace("{序号}", codeSeq);
                    }
                }
            }
            string orgId = "180000000000000";
            string transType = null;
            var smsResult = smsManager.SendSms(orgId, phone, smsContent, "system", transType,
                                               hostTrace.ToString(), Shop.SysParams.Default.SmsValidChannel);



            return JsonHelper.ObjectToJson(new
            {
                returnCode = "0000",
                returnMessage = "发送成功",
                smsCode = code,
                smsCodeSeq = codeSeq
            });
        }


        private string ProjectSearchApply(string idx, string idxname, string actno
            , string id, string mobile, string waterMark)
        {
            WebTool.WebTool.CheckerParametersIsNull("3264", "必填参数不能为空", idxname, actno
           , id, mobile, waterMark);

            /* if (!WebTool.WebTool.CheckIDCard(id))
             {
                 throw new TransactionException("1427", "身份证号不符合标准");
             }
             if (!WebTool.WebTool.ValidateMobile(mobile))
             {
                 throw new TransactionException("1427", "手机号不符合标准");
             }*/
            if (actno.Length < 16)
            {
                throw new TransactionException("1427", "银行卡不符合标准");
            }

            var strId = WebTool.WebTool.FromatUserId(id);
            var strMobile = WebTool.WebTool.FromatUserId(mobile);
            var strActNo = WebTool.WebTool.FromatUserId(actno);
            using (var p = new Processer())
            {

                //                var sql = new SqlBuilder();
                //                //校验验证码
                //                sql.AppendFormatAndParse(p, "SELECT DATA_VALUE FROM ULTAB_BO_SESSION_DATA a WHERE" +
                //                            " a.USERID = @USERID AND a.DATA_ID = @DATA_ID", mobile, validCodeName);
                //                var validInfo = p.ExecuteDataSet(sql);
                //                if (Checker.IsEmpty(validInfo) || waterMark != validInfo.PickValue("DATA_VALUE"))
                //                {
                //                    throw new TransactionException("3268", "验证码输入不正确");
                //                }
                //                //清除验证码
                //                sql.AppendFormatAndParse(p, @"DELETE FROM ULTAB_BO_SESSION_DATA 
                //WHERE USERID=@USERID
                //AND DATA_ID = @DATA_ID
                //", mobile, validCodeName);
                //                p.ExecuteNonQuery(sql);
                //                sql.Reset();

                string sessionKey = "ShopService:ZjjhReceipts:validcode_" + mobile;
                string volidCode = CacheHelper.Get(sessionKey).ToString();
                if (volidCode == null || Checker.IsEmpty(volidCode) || waterMark != volidCode)
                {
                    throw new TransactionException("3268", "验证码输入不正确");
                }

                var dr = new Dictionary<string, object>();
                try
                {
                    var actTraceNo = p.GetSequenceNext("UL_ACTIVITY_ID").ToString();
                    var now = DateTime.Now;
                    var date = now.ToString("yyyyMMddHHmmss");
                    var data = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                        //"         <entno>07901</entno>\n" +
                        "         <idx>{7}</idx>\n" +
                        "         <idxname>{8}</idxname>\n" +
                        "         <actno>{9}</actno>\n" +
                        //"         <idtype>{11}</idtype>\n" +
                        "         <id>{10}</id>\n" +
                        "         <mobile>{11}</mobile>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK05", date.Substring(0, 8), date.Substring(8), "cn", projectNo
                        , idx, idxname, actno, id, mobile);


                    var dataStr = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                        //"         <entno>{7}</entno>\n" +
                        "         <idx>{7}</idx>\n" +
                        "         <idxname>{8}</idxname>\n" +
                        "         <actno>{9}</actno>\n" +
                        //"         <idtype>{11}</idtype>\n" +
                        "         <id>{10}</id>\n" +
                        "         <mobile>{11}</mobile>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK05", date.Substring(0, 8), date.Substring(8), "cn", projectNo
                        , idx, idxname, strActNo, strId, strMobile);

                    Logger.LogInfo("调用4.5 代扣签约查询申请数据：" + dataStr);
                    var resultXml = Send(data);
                    //var resultXml = "<root> <head> <request_no>13970</request_no> <company_no>07901</company_no> <trcode>WXDK05</trcode> <retcode>0000</retcode> <retmsg>K_0000:交易成功</retmsg> <language>cn</language><txnodeid>07915</txnodeid> </head></root>";
                    Logger.LogInfo("4.5 代扣签约查询申请返回数据：" + resultXml);
                    if (Checker.IsEmpty(resultXml))
                    {
                        throw new TransactionException("0096", "第三方通讯异常");
                    }
                    var doc = new XmlDocument();
                    doc.LoadXml(resultXml);

                    var retcode = doc.GetElementsByTagName("retcode")[0].InnerText;
                    var retmsg = doc.GetElementsByTagName("retmsg")[0].InnerText;

                    if (retcode != "K_0000")
                    {
                        throw new TransactionException("5555", retmsg);
                    }

                    dr.Add("returnCode", "0000");
                    dr.Add("returnMessage", "交易成功");
                }
                catch (TransactionException ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", ex.ErrorCode);
                    dr.Add("returnMessage", ex.Message);
                }
                catch (Exception ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", "5555");
                    dr.Add("returnMessage", "系统提交失败,详询95533");
                }
                finally
                {
                    p.Dispose();
                }
                return JsonHelper.ObjectToJson(dr);
            }
        }

        private string ProjectSearchValid(string idx, string actno, string mobile, string validcode)
        {
            //entno 
            WebTool.WebTool.CheckerParametersIsNull("3264", "必填参数不能为空", actno, mobile);

            /*if (!WebTool.WebTool.ValidateMobile(mobile))
            {
                throw new TransactionException("1427", "手机号不符合标准");
            }*/
            if (actno.Length < 16)
            {
                throw new TransactionException("1427", "银行卡不符合标准");
            }

            var strMobile = WebTool.WebTool.FromatUserId(mobile);
            var strActNo = WebTool.WebTool.FromatUserId(actno);

            using (var p = new Processer())
            {
                var dr = new Dictionary<string, object>();
                try
                {
                    var actTraceNo = p.GetSequenceNext("UL_ACTIVITY_ID").ToString();
                    var now = DateTime.Now;
                    var date = now.ToString("yyyyMMddHHmmss");
                    var data = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                        //"         <entno>{7}</entno>\n" +
                        "         <idx>{7}</idx>\n" +
                        "         <actno>{8}</actno>\n" +
                        "         <mobile>{9}</mobile>\n" +
                        "         <validcode>{10}</validcode>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK06", date.Substring(0, 8), date.Substring(8), "cn", projectNo
                        , idx, actno, mobile, validcode);

                    var dataStr = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                        //"         <entno>{7}</entno>\n" +
                        "         <idx>{7}</idx>\n" +
                        "         <actno>{8}</actno>\n" +
                        "         <mobile>{9}</mobile>\n" +
                        "         <validcode>{10}</validcode>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK06", date.Substring(0, 8), date.Substring(8), "cn", projectNo
                        , idx, strActNo, strMobile, validcode);

                    Logger.LogInfo("调用代扣签约验证生效申请数据：" + dataStr);
                    var resultXml = Send(data);
                    //var resultXml = "<root> <head> <request_no>13971</request_no> <company_no>07901</company_no> <trcode>WXDK06</trcode> <retcode>0000</retcode> <retmsg>K_0000:交易成功</retmsg> <language>cn</language> <txnodeid>07915</txnodeid> </head> <body> <regcus> <entname>杭银消费</entname> <idx>181015153550042937</idx> <adddate>20181224</adddate> <entno>07901</entno> <stat>1</stat> </regcus> </body></root>";
                    Logger.LogInfo("代扣签约验证生效返回数据：" + resultXml);
                    if (Checker.IsEmpty(resultXml))
                    {
                        throw new TransactionException("0096", "第三方通讯异常");
                    }
                    var doc = new XmlDocument();
                    doc.LoadXml(resultXml);
                    var retcode = doc.GetElementsByTagName("retcode")[0].InnerText;
                    var retmsg = doc.GetElementsByTagName("retmsg")[0].InnerText;

                    var returnCode = "0000";
                    var returnMsg = "交易成功";


                    if (retcode != "K_0000")
                    {
                        returnCode = "5555";
                        returnMsg = retmsg;
                    }
                    else
                    {
                        var list = new List<object>();
                        var entinfo = doc.GetElementsByTagName("regcus");
                        for (var i = 0; i < entinfo.Count; i++)
                        {
                            var dr1 = new Dictionary<string, string>();
                            var entinfoVal = entinfo[i];
                            var entname = entinfoVal["entname"].InnerText;
                            var nowEntno = entinfoVal["entno"].InnerText;
                            var nowIdx = entinfoVal["idx"].InnerText;
                            var adddate = entinfoVal["adddate"].InnerText;

                            //0-	已申请  1-	已生效  2-	已撤销
                            var stat = entinfoVal["stat"].InnerText;
                            if (stat == "0" || stat == "1")
                            {
                                dr1.Add("entname", entname);
                                dr1.Add("entno", nowEntno);
                                dr1.Add("idx", nowIdx);
                                dr1.Add("adddate", adddate);
                                list.Add(dr1);
                            }
                        }
                        dr.Add("projectList", list);
                    }

                    if (retcode != "K_0000")
                    {
                        throw new TransactionException("5555", retmsg);
                    }

                    dr.Add("returnCode", returnCode);
                    dr.Add("returnMessage", returnMsg);
                }
                catch (TransactionException ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", ex.ErrorCode);
                    dr.Add("returnMessage", ex.Message);
                }
                catch (Exception ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", "5555");
                    dr.Add("returnMessage", "系统提交失败,详询95533");
                }
                finally
                {
                    p.Dispose();
                }
                return JsonHelper.ObjectToJson(dr);
            }
        }


        private string ProjectCancelApply(string entno, string idx, string actno, string mobile, string waterMark)
        {
            WebTool.WebTool.CheckerParametersIsNull("3264", "必填参数不能为空", entno, actno
         , mobile);

            /* if (!WebTool.WebTool.ValidateMobile(mobile))
             {
                 throw new TransactionException("1427", "手机号不符合标准");
             }*/
            if (actno.Length < 16)
            {
                throw new TransactionException("1427", "银行卡不符合标准");
            }

            var strMobile = WebTool.WebTool.FromatUserId(mobile);
            var strActNo = WebTool.WebTool.FromatUserId(actno);
            using (var p = new Processer())
            {

                if (Checker.IsEmpty(waterMark))
                {
                    throw new TransactionException("1427", "验证码不能为空");
                }

                //                var sql = new SqlBuilder();
                //                //校验验证码
                //                sql.AppendFormatAndParse(p, "SELECT DATA_VALUE FROM ULTAB_BO_SESSION_DATA a WHERE" +
                //                            " a.USERID = @USERID AND a.DATA_ID = @DATA_ID", mobile, validCodeName);
                //                var validInfo = p.ExecuteDataSet(sql);
                //                if (Checker.IsEmpty(validInfo) || waterMark != validInfo.PickValue("DATA_VALUE"))
                //                {
                //                    throw new TransactionException("3268", "验证码输入不正确");
                //                }
                //                //清除验证码
                //                sql.AppendFormatAndParse(p, @"DELETE FROM ULTAB_BO_SESSION_DATA 
                //WHERE USERID=@USERID
                //AND DATA_ID = @DATA_ID
                //", mobile, validCodeName);
                //                p.ExecuteNonQuery(sql);
                //                sql.Reset();

                string sessionKey = "ShopService:ZjjhReceipts:validcode_" + mobile;
                string volidCode = CacheHelper.Get(sessionKey).ToString();
                if (volidCode == null || Checker.IsEmpty(volidCode) || waterMark != volidCode)
                {
                    throw new TransactionException("3268", "验证码输入不正确");
                }
                var dr = new Dictionary<string, object>();
                try
                {
                    var actTraceNo = p.GetSequenceNext("UL_ACTIVITY_ID").ToString();
                    var now = DateTime.Now;
                    var date = now.ToString("yyyyMMddHHmmss");
                    var data = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                        "         <entno>{7}</entno>\n" +
                        "         <idx>{8}</idx>\n" +
                        "         <actno>{9}</actno>\n" +
                        "         <mobile>{10}</mobile>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK03", date.Substring(0, 8), date.Substring(8), "cn", projectNo
                        , entno, idx, actno, mobile);

                    var dataStr = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                        "         <entno>{7}</entno>\n" +
                        "         <idx>{8}</idx>\n" +
                        "         <actno>{9}</actno>\n" +
                        "         <mobile>{10}</mobile>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK03", date.Substring(0, 8), date.Substring(8), "cn", projectNo
                        , entno, idx, strActNo, strMobile);

                    Logger.LogInfo("调用建行代扣签约撤销申请数据：" + dataStr);
                    var resultXml = Send(data);
                    Logger.LogInfo("建行建行代扣签约撤销申请返回数据：" + resultXml);
                    if (Checker.IsEmpty(resultXml))
                    {
                        throw new TransactionException("0096", "第三方通讯异常");
                    }
                    var doc = new XmlDocument();
                    doc.LoadXml(resultXml);
                    var retcode = doc.GetElementsByTagName("retcode")[0].InnerText;
                    var retmsg = doc.GetElementsByTagName("retmsg")[0].InnerText;

                    if (retcode != "K_0000")
                    {
                        throw new TransactionException("5555", retmsg);
                    }

                    dr.Add("returnCode", "0000");
                    dr.Add("returnMessage", "交易成功");
                }
                catch (TransactionException ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", ex.ErrorCode);
                    dr.Add("returnMessage", ex.Message);
                }
                catch (Exception ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", "5555");
                    dr.Add("returnMessage", "系统提交失败,详询95533");
                }
                finally
                {
                    p.Dispose();
                }
                return JsonHelper.ObjectToJson(dr);
            }
        }

        private string ProjectCancelValid(string entno, string idx, string actno, string mobile, string validcode)
        {
            WebTool.WebTool.CheckerParametersIsNull("3264", "必填参数不能为空", entno, actno
         , mobile);

            /* if (!WebTool.WebTool.ValidateMobile(mobile))
             {
                 throw new TransactionException("1427", "手机号不符合标准");
             }*/
            if (actno.Length < 16)
            {
                throw new TransactionException("1427", "银行卡不符合标准");
            }

            var strMobile = WebTool.WebTool.FromatUserId(mobile);
            var strActNo = WebTool.WebTool.FromatUserId(actno);
            using (var p = new Processer())
            {
                var dr = new Dictionary<string, object>();
                try
                {
                    var actTraceNo = p.GetSequenceNext("UL_ACTIVITY_ID").ToString();
                    var now = DateTime.Now;
                    var date = now.ToString("yyyyMMddHHmmss");

                    var data = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                        "         <entno>{7}</entno>\n" +
                        "         <idx>{8}</idx>\n" +
                        "         <actno>{9}</actno>\n" +
                        "         <mobile>{10}</mobile>\n" +
                        "         <validcode>{11}</validcode>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK04", date.Substring(0, 8), date.Substring(8), "cn", projectNo
                        , entno, idx, actno, mobile, validcode);

                    var dataStr = string.Format(
                       "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                       "<root>\n" +
                       "   <head>\n" +
                       "      <request_no>{0}</request_no>\n" +
                       "      <company_no>{1}</company_no>\n" +
                       "      <trcode>{2}</trcode>\n" +
                       "      <trdate>{3}</trdate>\n" +
                       "      <trtime>{4}</trtime>\n" +
                       "      <language>{5}</language>\n" +
                       "      <txnodeid>{6}</txnodeid>\n" +
                       "   </head>\n" +
                       "   <body>\n" +
                       "         <entno>{7}</entno>\n" +
                       "         <idx>{8}</idx>\n" +
                       "         <actno>{9}</actno>\n" +
                       "         <mobile>{10}</mobile>\n" +
                       "         <validcode>{11}</validcode>\n" +
                       "   </body>\n" +
                       "</root>\n",
                       actTraceNo, companyNo, "WXDK04", date.Substring(0, 8), date.Substring(8), "cn", projectNo
                       , entno, idx, strActNo, strMobile, validcode);

                    Logger.LogInfo("调用代扣签约撤销验证生效申请数据：" + dataStr);
                    var resultXml = Send(data);

                    //var resultXml =
                    //  "<root> <head> <request_no>13961</request_no> <company_no>07901</company_no> <trcode>003010C832</trcode> <retcode>0000</retcode> <retmsg>K_0000:交易成功</retmsg> <language>cn</language> <txnodeid>07915</txnodeid> </head> <body> <cst_accno></cst_accno> <cst_accno_nm></cst_accno_nm> <bank_no></bank_no> <accno_stcd></accno_stcd> <acc_ap_dt></acc_ap_dt> <acc_op_dt></acc_op_dt> <acc_fl_dt></acc_fl_dt> </body></root>";
                    Logger.LogInfo("代扣签约撤销验证生效返回数据：" + resultXml);
                    if (Checker.IsEmpty(resultXml))
                    {
                        throw new TransactionException("0096", "第三方通讯异常");
                    }
                    var doc = new XmlDocument();
                    doc.LoadXml(resultXml);

                    var retcode = doc.GetElementsByTagName("retcode")[0].InnerText;
                    var retmsg = doc.GetElementsByTagName("retmsg")[0].InnerText;

                    if (retcode != "K_0000")
                    {
                        throw new TransactionException("5555", retmsg);
                    }

                    dr.Add("returnCode", "0000");
                    dr.Add("returnMessage", "交易成功");
                }
                catch (TransactionException ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", ex.ErrorCode);
                    dr.Add("returnMessage", ex.Message);
                }
                catch (Exception ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", "5555");
                    dr.Add("returnMessage", "系统提交失败,详询95533");
                }
                finally
                {
                    p.Dispose();
                }
                return JsonHelper.ObjectToJson(dr);
            }
        }


        private string ProjectApplyValid(string entno, string idx, string actno, string mobile, string validcode)
        {

            //todo 后续删掉
            //var result = "123";
            Logger.LogInfo("参数内容：entno:" + entno + " idx:" + idx + " actno:" + actno
                 + " mobile:" + mobile + " validcode:" + validcode);
            WebTool.WebTool.CheckerParametersIsNull("3264", "必填参数不能为空", entno, actno
          , mobile);

            /* if (!WebTool.WebTool.ValidateMobile(mobile))
             {
                 throw new TransactionException("1427", "手机号不符合标准");
             }*/
            if (actno.Length < 16)
            {
                throw new TransactionException("1427", "银行卡不符合标准");
            }

            var strMobile = WebTool.WebTool.FromatUserId(mobile);
            var strActNo = WebTool.WebTool.FromatUserId(actno);

            using (var p = new Processer())
            {
                var dr = new Dictionary<string, object>();
                try
                {
                    var actTraceNo = p.GetSequenceNext("UL_ACTIVITY_ID").ToString();
                    var now = DateTime.Now;
                    var date = now.ToString("yyyyMMddHHmmss");
                    var data = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                        "         <entno>{7}</entno>\n" +
                        "         <idx>{8}</idx>\n" +
                        "         <actno>{9}</actno>\n" +
                        "         <mobile>{10}</mobile>\n" +
                        "         <validcode>{11}</validcode>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK02", date.Substring(0, 8), date.Substring(8), "cn", projectNo
                        , entno, idx, actno, mobile, validcode);
                    var dataStr = string.Format(
                       "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                       "<root>\n" +
                       "   <head>\n" +
                       "      <request_no>{0}</request_no>\n" +
                       "      <company_no>{1}</company_no>\n" +
                       "      <trcode>{2}</trcode>\n" +
                       "      <trdate>{3}</trdate>\n" +
                       "      <trtime>{4}</trtime>\n" +
                       "      <language>{5}</language>\n" +
                       "      <txnodeid>{6}</txnodeid>\n" +
                       "   </head>\n" +
                       "   <body>\n" +
                       "         <entno>{7}</entno>\n" +
                       "         <idx>{8}</idx>\n" +
                       "         <actno>{9}</actno>\n" +
                       "         <mobile>{10}</mobile>\n" +
                       "         <validcode>{11}</validcode>\n" +
                       "   </body>\n" +
                       "</root>\n",
                       actTraceNo, companyNo, "WXDK02", date.Substring(0, 8), date.Substring(8), "cn", projectNo
                       , entno, idx, strActNo, strMobile, validcode);

                    Logger.LogInfo("调用代扣签约验证生效申请数据：" + dataStr);
                    var resultXml = Send(data);
                    Logger.LogInfo("代扣签约验证生效返回数据：" + resultXml);
                    if (Checker.IsEmpty(resultXml))
                    {
                        throw new TransactionException("0096", "第三方通讯异常");
                    }
                    var doc = new XmlDocument();
                    doc.LoadXml(resultXml);
                    var retcode = doc.GetElementsByTagName("retcode")[0].InnerText;
                    var retmsg = doc.GetElementsByTagName("retmsg")[0].InnerText;

                    if (retcode != "K_0000")
                    {
                        throw new TransactionException("5555", retmsg);
                    }

                    dr.Add("returnCode", "0000");
                    dr.Add("returnMessage", "交易成功");

                    //result = "{\"returnCode\":\"0000\",\"returnMessage\":\"交易成功\"}";
                }
                catch (TransactionException ex)
                {
                    Logger.LogException(ex);
                    /* dr.Add("returnCode", ex.ErrorCode);
                     dr.Add("returnMessage", ex.Message);*/
                    throw new TransactionException(ex.ErrorCode, ex.Message);
                }
                catch (Exception ex)
                {
                    Logger.LogException(ex);
                    throw new TransactionException("5555", "系统提交失败,详询95533");
                }
                finally
                {
                    p.Dispose();
                }
                return JsonHelper.ObjectToJson(dr);
                //return result;
            }
        }
        private string DeductMoneyProjectApply(string entno, string idx, string idxname, string actno
            , string id, string mobile, string waterMark)
        {
            //todo 后续删掉
            Logger.LogInfo("参数内容：entno:" + entno + " idx:" + idx + " idxname:" + idxname + " actno:" + actno
                + " id:" + id + " mobile:" + mobile + " waterMark:" + waterMark);
            WebTool.WebTool.CheckerParametersIsNull("3264", "必填参数不能为空", entno, idx, idxname, actno
            , id, mobile, waterMark);

            /*if (!WebTool.WebTool.CheckIDCard(id))
            {
                throw new TransactionException("1427", "身份证号不符合标准");
            }*/
            /* if (!WebTool.WebTool.ValidateMobile(mobile))
             {
                 throw new TransactionException("1427", "手机号不符合标准");
             }*/
            if (actno.Length < 16)
            {
                throw new TransactionException("1427", "银行卡不符合标准");
            }

            var strId = WebTool.WebTool.FromatUserId(id);
            var strMobile = WebTool.WebTool.FromatUserId(mobile);
            var strActNo = WebTool.WebTool.FromatUserId(actno);


            using (var p = new Processer())
            {
                if (Checker.IsEmpty(waterMark))
                {
                    throw new TransactionException("1427", "验证码不能为空");
                }


                //                var sql = new SqlBuilder();
                //                //校验验证码
                //                sql.AppendFormatAndParse(p, "SELECT DATA_VALUE FROM ULTAB_BO_SESSION_DATA a WHERE" +
                //                            " a.USERID = @USERID AND a.DATA_ID = @DATA_ID", mobile, validCodeName);
                //                var validInfo = p.ExecuteDataSet(sql);
                //                if (Checker.IsEmpty(validInfo) || waterMark != validInfo.PickValue("DATA_VALUE"))
                //                {
                //                    throw new TransactionException("3268", "验证码输入不正确");
                //                }
                //                //清除验证码
                //                sql.AppendFormatAndParse(p, @"DELETE FROM ULTAB_BO_SESSION_DATA 
                //WHERE USERID=@USERID
                //AND DATA_ID = @DATA_ID
                //", mobile, validCodeName);
                //                p.ExecuteNonQuery(sql);
                //                sql.Reset();

                string sessionKey = "ShopService:ZjjhReceipts:validcode_" + mobile;
                string volidCode = CacheHelper.Get(sessionKey).ToString();
                if (volidCode == null || Checker.IsEmpty(volidCode) || waterMark != volidCode)
                {
                    throw new TransactionException("3268", "验证码输入不正确");
                }

                var dr = new Dictionary<string, object>();
                var returnCode = "0000";
                var returnMsg = "交易成功";
                try
                {
                    var actTraceNo = p.GetSequenceNext("UL_ACTIVITY_ID").ToString();
                    var now = DateTime.Now;
                    var date = now.ToString("yyyyMMddHHmmss");
                    var nowDate = now.ToString("yyyyMMdd");
                    var nowTime = now.ToString("HHmmss");
                    var data = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                        "         <entno>{7}</entno>\n" +
                        "         <idx>{8}</idx>\n" +
                        "         <idxname>{9}</idxname>\n" +
                        "         <actno>{10}</actno>\n" +
                        // "         <idtype>{11}</idtype>\n" +
                        "         <id>{11}</id>\n" +
                        "         <mobile>{12}</mobile>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK01", date.Substring(0, 8), date.Substring(8), "cn", projectNo
                        , entno, idx, idxname, actno, id, mobile);

                    var dataStr = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                        "         <entno>{7}</entno>\n" +
                        "         <idx>{8}</idx>\n" +
                        "         <idxname>{9}</idxname>\n" +
                        "         <actno>{10}</actno>\n" +
                        // "         <idtype>{11}</idtype>\n" +
                        "         <id>{11}</id>\n" +
                        "         <mobile>{12}</mobile>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK01", date.Substring(0, 8), date.Substring(8), "cn", projectNo
                        , entno, idx, idxname, strActNo, strId, strMobile);

                    Logger.LogInfo("调用建行代扣签约申请数据：" + dataStr);
                    var resultXml = Send(data);
                    var rsp = DateTime.Now;
                    var rspDate = rsp.ToString("yyyyMMdd");
                    var rspTime = rsp.ToString("HHmmss");
                    Logger.LogInfo("建行代扣签约申请返回数据：" + resultXml);
                    if (Checker.IsEmpty(resultXml))
                    {
                        throw new TransactionException("0096", "第三方通讯异常");
                    }
                    var doc = new XmlDocument();
                    doc.LoadXml(resultXml);
                    var retcode = doc.GetElementsByTagName("retcode")[0].InnerText;
                    var retmsg = doc.GetElementsByTagName("retmsg")[0].InnerText;

                    if (retcode != "K_0000")
                    {
                        returnCode = "5555";
                        returnMsg = retmsg;
                    }

                    SaveCallBackLs(null, null, null, null, "WXDK01", "sha256", null, null, null, null, null
                      , actTraceNo, retcode, retmsg, dataStr, resultXml, null, null, null, nowDate, nowTime
                      , rspDate, rspTime, "0000", "交易成功", null, null);


                    dr.Add("returnCode", returnCode);
                    dr.Add("returnMessage", returnMsg);
                }
                catch (TransactionException ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", ex.ErrorCode);
                    dr.Add("returnMessage", ex.Message);
                }
                catch (Exception ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", "5555");
                    dr.Add("returnMessage", "系统提交失败,详询95533");
                }
                finally
                {
                    p.Dispose();
                }
                return JsonHelper.ObjectToJson(dr);
            }
        }

        public string Test(string name, string cardNo, string idNo, string mobile)
        {
            var returnCode = "0000";
            var returnMsg = "交易成功";

            using (var p = new Processer())
            {
                var dr = new Dictionary<string, object>();
                try
                {
                    var actTraceNo = p.GetSequenceNext("UL_ACTIVITY_ID").ToString();
                    var now = DateTime.Now;
                    var date = now.ToString("yyyyMMddHHmmss");
                    string data = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                          "   <actno>{7}</actno>\n" +
                          "   <id>{8}</id>\n" +
                          "   <idtype>1010</idtype>\n" +
                          "   <idx>181015153550042937</idx>\n" +
                          "   <idxname>{9}</idxname>\n" +
                          "   <mobile>{10}</mobile>\n" +
                          "   <TransNo>181015153550042937</TransNo>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK01", date.Substring(0, 8), date.Substring(8), "cn"
                        , projectNo, cardNo, idNo, name, mobile);

                    Logger.LogInfo("调用建行查询可扣款地区数据：" + data);
                    var resultXml = Send(data);//"101.71.36.85", 33021,
                    Logger.LogInfo("建行查询可扣款地区返回数据：" + resultXml);
                    if (Checker.IsEmpty(resultXml))
                    {
                        throw new TransactionException("0096", "第三方通讯异常");
                    }
                    var doc = new XmlDocument();
                    doc.LoadXml(resultXml);
                    var retcode = doc.GetElementsByTagName("retcode")[0].InnerText;
                    var retmsg = doc.GetElementsByTagName("retmsg")[0].InnerText;

                    if (retcode != "K_0000")
                    {
                        returnCode = "5555";
                        returnMsg = retmsg;
                    }
                    else
                    {
                        var prodarea = doc.GetElementsByTagName("prodarea")[0].InnerText;

                        prodarea = prodarea.PadLeft(10, '0');
                        var list = new List<object>();
                        var dr1 = new Dictionary<string, string>();
                        //1杭州 2温州 3嘉兴 4湖州 5绍兴 6台州 7金华 8衢州 9丽水 10舟山
                        list = GetCity(0, "杭州", prodarea, list);
                        list = GetCity(1, "温州", prodarea, list);
                        list = GetCity(2, "嘉兴", prodarea, list);
                        list = GetCity(3, "湖州", prodarea, list);
                        list = GetCity(4, "绍兴", prodarea, list);
                        list = GetCity(5, "台州", prodarea, list);
                        list = GetCity(6, "金华", prodarea, list);
                        list = GetCity(7, "衢州", prodarea, list);
                        list = GetCity(8, "丽水", prodarea, list);
                        list = GetCity(9, "舟山", prodarea, list);
                        //list.Add(dr1);
                        dr.Add("cityList", list);
                    }

                    dr.Add("returnCode", returnCode);
                    dr.Add("returnMessage", returnMsg);
                }
                catch (TransactionException ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", ex.ErrorCode);
                    dr.Add("returnMessage", ex.Message);
                }
                catch (Exception ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", "5555");
                    dr.Add("returnMessage", "系统提交失败,详询95533");
                }
                finally
                {
                    p.Dispose();
                }
                return JsonHelper.ObjectToJson(dr);
            }
        }

        public string QueryDeductMoneyArea()
        {
            /*
                        var testDr = new Dictionary<string, object>();
                        var testProdarea = "1011001110";
                        testProdarea = testProdarea.PadLeft(10, '0');
                        var list1 = new List<object>();

                        //1杭州 2温州 3嘉兴 4湖州 5绍兴 6台州 7金华 8衢州 9丽水 10舟山
                        list1 = GetCity(1, "杭州", testProdarea, list1);
                        list1 = GetCity(2, "温州", testProdarea, list1);
                        list1 = GetCity(3, "嘉兴", testProdarea, list1);
                        list1 = GetCity(4, "湖州", testProdarea, list1);
                        list1 = GetCity(5, "绍兴", testProdarea, list1);
                        list1 = GetCity(6, "台州", testProdarea, list1);
                        list1 = GetCity(7, "金华", testProdarea, list1);
                        list1 = GetCity(8, "衢州", testProdarea, list1);
                        list1 = GetCity(9, "丽水", testProdarea, list1);
                        list1 = GetCity(10, "舟山", testProdarea, list1);
                        //list1.Add(dr2);
                        testDr.Add("returnCode", "0000");
                        testDr.Add("returnMessage", "交易成功");
                        testDr.Add("cityList", list1);
                        return JsonHelper.ObjectToJson(testDr);
                        */

            var returnCode = "0000";
            var returnMsg = "交易成功";
            using (var p = new Processer())
            {
                var dr = new Dictionary<string, object>();
                try
                {
                    var actTraceNo = p.GetSequenceNext("UL_ACTIVITY_ID").ToString();
                    var now = DateTime.Now;
                    var date = now.ToString("yyyyMMddHHmmss");
                    string data = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK07", date.Substring(0, 8), date.Substring(8), "cn", projectNo);

                    Logger.LogInfo("调用建行查询可扣款地区数据：" + data);
                    var resultXml = Send(data);//"101.71.36.85", 33021,
                    Logger.LogInfo("建行查询可扣款地区返回数据：" + resultXml);
                    if (Checker.IsEmpty(resultXml))
                    {
                        throw new TransactionException("0096", "第三方通讯异常");
                    }
                    var doc = new XmlDocument();
                    doc.LoadXml(resultXml);
                    var retcode = doc.GetElementsByTagName("retcode")[0].InnerText;
                    var retmsg = doc.GetElementsByTagName("retmsg")[0].InnerText;

                    if (retcode != "K_0000")
                    {
                        returnCode = "5555";
                        returnMsg = retmsg;
                        Logger.LogInfo("接口返回错误：test " + retmsg);

                    }
                    else
                    {
                        var prodarea = doc.GetElementsByTagName("prodarea")[0].InnerText;
                        Logger.LogInfo(prodarea);
                        prodarea = prodarea.PadLeft(10, '0');
                        var list = new List<object>();
                        //var dr1 = new Dictionary<string, string>();
                        //1杭州 2温州 3嘉兴 4湖州 5绍兴 6台州 7金华 8衢州 9丽水 10舟山
                        list = GetCity(1, "杭州", prodarea, list);
                        list = GetCity(2, "温州", prodarea, list);
                        list = GetCity(3, "嘉兴", prodarea, list);
                        list = GetCity(4, "湖州", prodarea, list);
                        list = GetCity(5, "绍兴", prodarea, list);
                        list = GetCity(6, "台州", prodarea, list);
                        list = GetCity(7, "金华", prodarea, list);
                        list = GetCity(8, "衢州", prodarea, list);
                        list = GetCity(9, "丽水", prodarea, list);
                        list = GetCity(10, "舟山", prodarea, list);
                        //list.Add(dr1);
                        dr.Add("cityList", list);
                    }

                    dr.Add("returnCode", returnCode);
                    dr.Add("returnMessage", returnMsg);
                }
                catch (TransactionException ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", ex.ErrorCode);
                    dr.Add("returnMessage", ex.Message);
                }
                catch (Exception ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", "5555");
                    dr.Add("returnMessage", "系统提交失败,详询95533");
                }
                finally
                {
                    p.Dispose();
                }
                return JsonHelper.ObjectToJson(dr);
            }
        }

        public string QueryDeductMoneyProject(string areaId)
        {

            Logger.LogInfo("QueryDeductMoneyProject 参数是:" + areaId);
            /*
            var testDr = new Dictionary<string, object>();
            var zimuDic = new Dictionary<string, object>();
            var list1 = new List<object>();
            // var list2 = new List<object>();
            for (var j = 0; j < 5; j++)
            {
                var dr2 = new Dictionary<string, string>();
                dr2.Add("entname", "电费" + j.ToString());
                dr2.Add("entno", "02123");
                dr2.Add("reqidx", "1");//+ j.ToString()
                //dr2.Add("reqidx", "222222" + j.ToString());
                //dr2.Add("idxname", "333333" + j.ToString());
                dr2.Add("idxname", "cshi");
                list1.Add(dr2);
            }


            //zimuDic.Add("w", list1);
            //zimuDic.Add("s", list1);#1#

            testDr.Add("returnCode", "0000");
            testDr.Add("returnMessage", "交易成功");
            testDr.Add("projectList", list1);
            return JsonHelper.ObjectToJson(testDr);
            
        */

            WebTool.WebTool.CheckerParametersIsNull("3264", "必填参数不能为空", areaId);

            var returnCode = "0000";
            var returnMsg = "交易成功";
            using (var p = new Processer())
            {
                var dr = new Dictionary<string, object>();
                try
                {
                    var actTraceNo = p.GetSequenceNext("UL_ACTIVITY_ID").ToString();
                    var now = DateTime.Now;
                    var date = now.ToString("yyyyMMddHHmmss");
                    string data = string.Format(
                        "<?xml version=\"1.0\" encoding=\"GBK\"?>\n" +
                        "<root>\n" +
                        "   <head>\n" +
                        "      <request_no>{0}</request_no>\n" +
                        "      <company_no>{1}</company_no>\n" +
                        "      <trcode>{2}</trcode>\n" +
                        "      <trdate>{3}</trdate>\n" +
                        "      <trtime>{4}</trtime>\n" +
                        "      <language>{5}</language>\n" +
                        "      <txnodeid>{6}</txnodeid>\n" +
                        "   </head>\n" +
                        "   <body>\n" +
                        "         <area>{7}</area>\n" +
                        "   </body>\n" +
                        "</root>\n",
                        actTraceNo, companyNo, "WXDK08", date.Substring(0, 8), date.Substring(8), "cn", projectNo, areaId);

                    Logger.LogInfo("调用建行查询代扣项目数据：" + data);
                    var resultXml = Send(data);
                    Logger.LogInfo("建行查询代扣项目返回数据：" + resultXml);
                    if (Checker.IsEmpty(resultXml))
                    {
                        throw new TransactionException("0096", "第三方通讯异常");
                    }
                    var doc = new XmlDocument();
                    doc.LoadXml(resultXml);
                    var retcode = doc.GetElementsByTagName("retcode")[0].InnerText;
                    var retmsg = doc.GetElementsByTagName("retmsg")[0].InnerText;

                    if (retcode != "K_0000")
                    {
                        returnCode = "5555";
                        returnMsg = retmsg;
                    }
                    else
                    {
                        var list = new List<object>();
                        var entinfo = doc.GetElementsByTagName("entinfo");
                        for (var i = 0; i < entinfo.Count; i++)
                        {
                            var dr1 = new Dictionary<string, string>();
                            var entinfoVal = entinfo[i];
                            var entname = entinfoVal["entname"].InnerText;
                            var entno = entinfoVal["entno"].InnerText;
                            var reqidx = entinfoVal["reqidx"].InnerText;
                            var idxname = entinfoVal["idxname"].InnerText;
                            dr1.Add("entname", entname);
                            dr1.Add("entno", entno);
                            dr1.Add("reqidx", reqidx);
                            dr1.Add("idxname", idxname);
                            list.Add(dr1);
                        }
                        dr.Add("projectList", list);
                    }
                    dr.Add("returnCode", returnCode);
                    dr.Add("returnMessage", returnMsg);
                }
                catch (TransactionException ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", ex.ErrorCode);
                    dr.Add("returnMessage", ex.Message);
                }
                catch (Exception ex)
                {
                    Logger.LogException(ex);
                    dr.Add("returnCode", "5555");
                    dr.Add("returnMessage", "系统提交失败,详询95533");
                }
                finally
                {
                    p.Dispose();
                }
                return JsonHelper.ObjectToJson(dr);
            }
        }

        private List<object> GetCity(int i, string city, string prodarea, List<object> list)
        {
            var dr = new Dictionary<string, string>();
            if (prodarea.Substring(i, 1) == "1")
            {
                dr.Add("cityVal", i.ToString());
                dr.Add("cityName", city);
                list.Add(dr);
            }
            return list;
        }

        public string Send(String data)
        {
            try
            {
                Logger.LogInfo(string.Format("Send计算: {0},{1},{2}", host, port, data));
                var reqU = data + "\0";
                var mac = new byte[48];
                byte[] dataByte = Encoding.GetEncoding("GBK").GetBytes(reqU);
                var length = dataByte.Length;
                var ret = BhsService.sha256(dataByte, length, mac);
                Logger.LogInfo(string.Format("sha256计算: {0}", ret));
                if (ret == 0)
                {
                    IntPtr ptr = IntPtr.Zero;
                    Int64 recvlen = 1024;
                    var recvfilename = new byte[1024];
                    int ret2 = BhsService.dataexch(host, port, "exchbankfile",
                                                   Encoding.GetEncoding("GBK").GetString(mac),
                                                   dataByte, length, "", "b", ref ptr,
                                                   ref recvlen, recvfilename, 30);
                    Logger.LogInfo(string.Format("dataexch计算: {0}", ret2));
                    if (ret2 == 0)
                    {
                        string ptrs = Marshal.PtrToStringAnsi(ptr);
                        Logger.LogInfo(string.Format("dataexch计算: {0}", ptrs));
                        return ptrs;
                    }
                    throw new TransactionException("0096", "第三方通讯异常");
                }
            }
            catch (Exception ex)
            {
                Logger.LogInfo("send 出错：" + ex);
                Logger.LogException(ex);
            }
            return "";
        }

        public class BhsService
        {
            [DllImport(DllAddr)]
            public static extern int dataexch(String ip, int port, String transtype, String mac, byte[] send,
                int snedlen, String sendfilename, String mode,
                ref IntPtr receve, ref Int64 recvlen, byte[] recvfilename, int timeout);

            [DllImport(DllAddr)]
            public static extern int sha256(byte[] data, int size, byte[] result);

        }



        public int SaveCallBackLs(string orgId, string mercId, string orderId, string shoptraceNo, string transType
            , string signMethod, string signature, string backUrl, string amount, string transStatus, string hostTraceNo
            , string traceNo, string retCode, string retMsg, string reqMsg, string rspMsg, string reqencryptionMsg
            , string rspencryptionMsg, string reserved, string reqDate, string reqTime, string rspDate, string rspTime
            , string umqRetcode, string umqRetmsg, string mercTranstype, string umqTranstype)
        {
            using (var p = new Processer())
            {
                var sql = new SqlBuilder();
                sql.Append(@"INSERT INTO ULTAB_FO_CALLBACKTRANSLS (ORG_ID, MERC_ID, ORDERID, SHOPTRACENO, TRANSTYPE, SIGNMETHOD, SIGNATURE, BACKURL, AMOUNT, TRANSSTATUS, HOSTTRACENO, TRACENO, RETCODE, RETMSG, REQMSG, RSPMSG, REQENCRYPTIONMSG, RSPENCRYPTIONMSG, RESERVED, REQDATE, REQTIME, RSPDATE, RSPTIME, UMQRETCODE, UMQRETMSG, MERCTRANSTYPE, UMQTRANSTYPE)
VALUES (@l_org_id, @l_merc_id, @l_orderid, @l_shoptraceno, @l_transtype, @l_signmethod, @l_signature, @l_backurl, @l_amount, @l_transstatus, @l_hosttraceno, @l_traceno, @l_retcode, @l_retmsg, @l_reqmsg, @l_rspmsg, @l_reqencryptionmsg, @l_rspencryptionmsg, @l_reserved, @l_reqdate, @l_reqtime, @l_rspdate, @l_rsptime, @l_umqretcode, @l_umqretmsg, @l_merctranstype, @l_umqtranstype)");
                p.AddDataParameter("@l_org_id", orgId);
                p.AddDataParameter("@l_merc_id", mercId);
                p.AddDataParameter("@l_orderid", orderId);
                p.AddDataParameter("@l_shoptraceno", shoptraceNo);
                p.AddDataParameter("@l_transtype", transType);
                p.AddDataParameter("@l_signmethod", signMethod);
                p.AddDataParameter("@l_backurl", backUrl);
                p.AddDataParameter("@l_amount", amount);
                p.AddDataParameter("@l_transstatus", transStatus);
                p.AddDataParameter("@l_hosttraceno", hostTraceNo);
                p.AddDataParameter("@l_traceno", traceNo);
                p.AddDataParameter("@l_retcode", retCode);
                p.AddDataParameter("@l_retmsg", retMsg);
                p.AddDataParameter("@l_reqmsg", reqMsg);
                p.AddDataParameter("@l_rspmsg", rspMsg);
                p.AddDataParameter("@l_reqencryptionmsg", reqencryptionMsg);
                p.AddDataParameter("@l_rspencryptionmsg", rspencryptionMsg);
                p.AddDataParameter("@l_reserved", reserved);
                p.AddDataParameter("@l_reqdate", reqDate);
                p.AddDataParameter("@l_reqtime", reqTime);
                p.AddDataParameter("@l_rspdate", rspDate);
                p.AddDataParameter("@l_rsptime", rspTime);
                p.AddDataParameter("@l_umqretcode", umqRetcode);
                p.AddDataParameter("@l_umqretmsg", umqRetmsg);
                p.AddDataParameter("@l_merctranstype", mercTranstype);
                p.AddDataParameter("@l_umqtranstype", umqTranstype);
                var upNum = p.ExecuteNonQuery(sql);
                Logger.LogInfo(p.LastSQL);
                p.Commit();
                return upNum;
            }
        }


        #region 生成GetValidCode

        private const string validCodeName = "__VALID_CODE__";
        /// <summary>
        /// 生成GetValidCode
        /// </summary>
        /// <param name="mobile"></param>
        /// <returns></returns>
        private string GetValidCode(string mobile)
        {
            Logger.LogInfo("GetValidCode：" + mobile);
            WebTool.WebTool.CheckerParametersIsNull("3264", "必填参数不能为空", mobile);
            string sessionKey = "ShopService:ZjjhReceipts:validcode_" + mobile;
            var r = GetRandom();
            var rImg = CreateImage(r);
            var now = DateTime.Now;
            var nowTime = now.ToString("yyyyMMddHHmmssfff");
            //var sessionId = "zjjhSign";//浙江建行簽約
            CacheHelper.Insert(sessionKey, r, 120);


            //            var sql = new SqlBuilder();
            //            using (var p = new Processer())
            //            {
            //                sql.AppendFormatAndParse(p, @"UPDATE ULTAB_BO_SESSION_DATA 
            //SET LASTOPERTIME=@LASTOPERTIME,
            //DATA_VALUE = @DATA_VALUE
            //WHERE USERID=@USERID
            //AND DATA_ID = @DATA_ID
            //", nowTime, r, mobile, validCodeName);
            //                var result = p.ExecuteNonQuery(sql);

            //                if (result <= 0)
            //                {
            //                    sql.AppendFormatAndParse(p, @"INSERT INTO ULTAB_BO_SESSION_DATA (SESSIONID, FIRSTOPERTIME, LASTOPERTIME,
            //USERID, DATA_ID, DATA_VALUE)
            //VALUES ( @SESSIONID,@FIRSTOPERTIME, @LASTOPERTIME, @USERID, @DATA_ID, @DATA_VALUE)",
            //                     sessionId, nowTime, nowTime, mobile, validCodeName, r);
            //                    p.ExecuteNonQuery(sql);
            //                }
            //            }

            return JsonHelper.ObjectToJson(new
            {
                returnCode = "0000",
                returnMessage = "交易成功",
                code = rImg
            });

        }

        private string GetRandom()
        {
            var r = new Random();
            return r.Next(0000, 9999).ToString("0000");
        }

        private string CreateImage(string checkCode, string backColor = null)
        {
            var iwidth = checkCode.Length * 15;
            var image = new Bitmap(iwidth, 25);
            var g = Graphics.FromImage(image);
            var graphicBackColor = Checker.IsEmpty(backColor) ? Color.White : ColorTranslator.FromHtml("#" + backColor);
            g.Clear(graphicBackColor);
            //定义颜色
            Color[] c = { Color.Black, Color.Red, Color.DarkBlue, Color.Green, Color.DarkGray, Color.Brown, Color.DarkCyan, Color.Purple };
            //定义字体            
            string[] font = { "Verdana", "Microsoft Sans Serif", "Comic Sans MS", "Arial", "宋体" };
            var rand = new Random();
            //随机输出噪点
            for (int i = 0; i < 50; i++)
            {
                int x = rand.Next(image.Width);
                int y = rand.Next(image.Height);
                g.DrawRectangle(new Pen(Color.LightGray, 0), x, y, 1, 1);
            }

            //输出不同字体和颜色的验证码字符
            for (var i = 0; i < checkCode.Length; i++)
            {
                var cindex = rand.Next(7);
                var findex = rand.Next(5);

                var f = new Font(font[findex], 11, FontStyle.Bold);
                var b = new SolidBrush(c[cindex]);
                var ii = 4;
                if ((i + 1) % 2 == 0)
                {
                    ii = 2;
                }
                g.DrawString(checkCode.Substring(i, 1), f, b, 3 + (i * 12), ii);
            }
            //画一个边框
            //g.DrawRectangle(new Pen(Color.Black, 0), 0, 0, image.Width - 1, image.Height - 1);

            //输出到浏览器
            using (var ms = new MemoryStream())
            {
                image.Save(ms, System.Drawing.Imaging.ImageFormat.Jpeg);
                var result = Convert.ToBase64String(ms.ToArray());
                g.Dispose();
                image.Dispose();
                return result;
            }
        }

        #endregion

        public bool IsReusable
        {
            get
            {
                return false;
            }
        }
    }
}
