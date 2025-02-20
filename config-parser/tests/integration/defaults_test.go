// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package integration_test

import (
	"bytes"
	"testing"

	parser "github.com/haproxytech/client-native/v5/config-parser"
	"github.com/haproxytech/client-native/v5/config-parser/options"
)

func TestWholeConfigsSectionsDefaults(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Name, Config string
	}{
		{"defaults_balancehashreqcookieclientid", defaults_balancehashreqcookieclientid},
		{"defaults_balancehashreqhdripxforwardedfor", defaults_balancehashreqhdripxforwardedfor},
		{"defaults_balancehashreqhdripxforwardedfor_", defaults_balancehashreqhdripxforwardedfor_},
		{"defaults_balancehdrhdrName", defaults_balancehdrhdrName},
		{"defaults_balancehdrhdrNameusedomainonly", defaults_balancehdrhdrNameusedomainonly},
		{"defaults_balancerandom", defaults_balancerandom},
		{"defaults_balancerandom15", defaults_balancerandom15},
		{"defaults_balancerdpcookie", defaults_balancerdpcookie},
		{"defaults_balancerdpcookiesomething", defaults_balancerdpcookiesomething},
		{"defaults_balanceroundrobin", defaults_balanceroundrobin},
		{"defaults_balanceuri", defaults_balanceuri},
		{"defaults_balanceuridepth8", defaults_balanceuridepth8},
		{"defaults_balanceuridepth8len12whole", defaults_balanceuridepth8len12whole},
		{"defaults_balanceuridepth8whole", defaults_balanceuridepth8whole},
		{"defaults_balanceurilen12", defaults_balanceurilen12},
		{"defaults_balanceuriwhole", defaults_balanceuriwhole},
		{"defaults_balanceurlparam", defaults_balanceurlparam},
		{"defaults_balanceurlparamcheckpost10", defaults_balanceurlparamcheckpost10},
		{"defaults_balanceurlparamcheckpost10maxwai", defaults_balanceurlparamcheckpost10maxwai},
		{"defaults_balanceurlparamsessionid", defaults_balanceurlparamsessionid},
		{"defaults_balanceurlparamsessionidcheckpos", defaults_balanceurlparamsessionidcheckpos},
		{"defaults_cookiemyCookieattrSameSiteStrict", defaults_cookiemyCookieattrSameSiteStrict},
		{"defaults_cookiemyCookiedomaindom1domaindo", defaults_cookiemyCookiedomaindom1domaindo},
		{"defaults_cookiemyCookiedomaindom1domaindo_", defaults_cookiemyCookiedomaindom1domaindo_},
		{"defaults_cookiemyCookiedomaindom1indirect", defaults_cookiemyCookiedomaindom1indirect},
		{"defaults_cookiemyCookieindirectmaxidle10", defaults_cookiemyCookieindirectmaxidle10},
		{"defaults_cookiemyCookieindirectmaxidle10m", defaults_cookiemyCookieindirectmaxidle10m},
		{"defaults_cookiemyCookieindirectmaxlife10", defaults_cookiemyCookieindirectmaxlife10},
		{"defaults_cookietest", defaults_cookietest},
		{"defaults_defaultserveraddr1", defaults_defaultserveraddr1},
		{"defaults_defaultserveraddr127001", defaults_defaultserveraddr127001},
		{"defaults_defaultserveragentaddr127001", defaults_defaultserveragentaddr127001},
		{"defaults_defaultserveragentaddrsitecom", defaults_defaultserveragentaddrsitecom},
		{"defaults_defaultserveragentcheck", defaults_defaultserveragentcheck},
		{"defaults_defaultserveragentinter1000ms", defaults_defaultserveragentinter1000ms},
		{"defaults_defaultserveragentport1", defaults_defaultserveragentport1},
		{"defaults_defaultserveragentport65535", defaults_defaultserveragentport65535},
		{"defaults_defaultserveragentsendname", defaults_defaultserveragentsendname},
		{"defaults_defaultserverallow0rtt", defaults_defaultserverallow0rtt},
		{"defaults_defaultserveralpnh2", defaults_defaultserveralpnh2},
		{"defaults_defaultserveralpnh2http11", defaults_defaultserveralpnh2http11},
		{"defaults_defaultserveralpnhttp11", defaults_defaultserveralpnhttp11},
		{"defaults_defaultserverbackup", defaults_defaultserverbackup},
		{"defaults_defaultservercafilecertcrt", defaults_defaultservercafilecertcrt},
		{"defaults_defaultservercheck", defaults_defaultservercheck},
		{"defaults_defaultservercheckalpnhttp10", defaults_defaultservercheckalpnhttp10},
		{"defaults_defaultservercheckalpnhttp11http", defaults_defaultservercheckalpnhttp11http},
		{"defaults_defaultservercheckprotoh2", defaults_defaultservercheckprotoh2},
		{"defaults_defaultserverchecksendproxy", defaults_defaultserverchecksendproxy},
		{"defaults_defaultservercheckssl", defaults_defaultservercheckssl},
		{"defaults_defaultservercheckviasocks4", defaults_defaultservercheckviasocks4},
		{"defaults_defaultserverciphersECDHEECDSACH", defaults_defaultserverciphersECDHEECDSACH},
		{"defaults_defaultserverciphersECDHERSAAES1", defaults_defaultserverciphersECDHERSAAES1},
		{"defaults_defaultserverciphersuitesECDHEEC", defaults_defaultserverciphersuitesECDHEEC},
		{"defaults_defaultservercookievalue", defaults_defaultservercookievalue},
		{"defaults_defaultservercrlfilefilepem", defaults_defaultservercrlfilefilepem},
		{"defaults_defaultservercrtcertpem", defaults_defaultservercrtcertpem},
		{"defaults_defaultserverdisabled", defaults_defaultserverdisabled},
		{"defaults_defaultserverdowninter3500ms", defaults_defaultserverdowninter3500ms},
		{"defaults_defaultserverenabled", defaults_defaultserverenabled},
		{"defaults_defaultservererrorlimit50", defaults_defaultservererrorlimit50},
		{"defaults_defaultserverfall1rise2inter3spo", defaults_defaultserverfall1rise2inter3spo},
		{"defaults_defaultserverfall30", defaults_defaultserverfall30},
		{"defaults_defaultserverfastinter2500ms", defaults_defaultserverfastinter2500ms},
		{"defaults_defaultserverfastinterunknown", defaults_defaultserverfastinterunknown},
		{"defaults_defaultserverforcesslv3", defaults_defaultserverforcesslv3},
		{"defaults_defaultserverforcetlsv10", defaults_defaultserverforcetlsv10},
		{"defaults_defaultserverforcetlsv11", defaults_defaultserverforcetlsv11},
		{"defaults_defaultserverforcetlsv12", defaults_defaultserverforcetlsv12},
		{"defaults_defaultserverforcetlsv13", defaults_defaultserverforcetlsv13},
		{"defaults_defaultserverinitaddrlastlibcnon", defaults_defaultserverinitaddrlastlibcnon},
		{"defaults_defaultserverinitaddrlastlibcnon_", defaults_defaultserverinitaddrlastlibcnon_},
		{"defaults_defaultserverinter1000weight13", defaults_defaultserverinter1000weight13},
		{"defaults_defaultserverinter1500ms", defaults_defaultserverinter1500ms},
		{"defaults_defaultserverlogbufsize10", defaults_defaultserverlogbufsize10},
		{"defaults_defaultserverlogprotolegacy", defaults_defaultserverlogprotolegacy},
		{"defaults_defaultserverlogprotooctetcount", defaults_defaultserverlogprotooctetcount},
		{"defaults_defaultservermaxconn1", defaults_defaultservermaxconn1},
		{"defaults_defaultservermaxconn50", defaults_defaultservermaxconn50},
		{"defaults_defaultservermaxqueue0", defaults_defaultservermaxqueue0},
		{"defaults_defaultservermaxqueue1000", defaults_defaultservermaxqueue1000},
		{"defaults_defaultservermaxreuse0", defaults_defaultservermaxreuse0},
		{"defaults_defaultservermaxreuse1", defaults_defaultservermaxreuse1},
		{"defaults_defaultservermaxreuse1_", defaults_defaultservermaxreuse1_},
		{"defaults_defaultserverminconn1", defaults_defaultserverminconn1},
		{"defaults_defaultserverminconn50", defaults_defaultserverminconn50},
		{"defaults_defaultservernamespacetest", defaults_defaultservernamespacetest},
		{"defaults_defaultservernoagentcheck", defaults_defaultservernoagentcheck},
		{"defaults_defaultservernobackup", defaults_defaultservernobackup},
		{"defaults_defaultservernocheck", defaults_defaultservernocheck},
		{"defaults_defaultservernocheckssl", defaults_defaultservernocheckssl},
		{"defaults_defaultservernonstick", defaults_defaultservernonstick},
		{"defaults_defaultservernosendproxyv2", defaults_defaultservernosendproxyv2},
		{"defaults_defaultservernosendproxyv2ssl", defaults_defaultservernosendproxyv2ssl},
		{"defaults_defaultservernosendproxyv2sslcn", defaults_defaultservernosendproxyv2sslcn},
		{"defaults_defaultservernossl", defaults_defaultservernossl},
		{"defaults_defaultservernosslreuse", defaults_defaultservernosslreuse},
		{"defaults_defaultservernosslv3", defaults_defaultservernosslv3},
		{"defaults_defaultservernotfo", defaults_defaultservernotfo},
		{"defaults_defaultservernotlstickets", defaults_defaultservernotlstickets},
		{"defaults_defaultservernotlsv10", defaults_defaultservernotlsv10},
		{"defaults_defaultservernotlsv11", defaults_defaultservernotlsv11},
		{"defaults_defaultservernotlsv12", defaults_defaultservernotlsv12},
		{"defaults_defaultservernotlsv13", defaults_defaultservernotlsv13},
		{"defaults_defaultservernoverifyhost", defaults_defaultservernoverifyhost},
		{"defaults_defaultservernpnhttp11http10", defaults_defaultservernpnhttp11http10},
		{"defaults_defaultserverobservelayer4", defaults_defaultserverobservelayer4},
		{"defaults_defaultserverobservelayer7", defaults_defaultserverobservelayer7},
		{"defaults_defaultserveronerrorfailcheck", defaults_defaultserveronerrorfailcheck},
		{"defaults_defaultserveronerrorfastinter", defaults_defaultserveronerrorfastinter},
		{"defaults_defaultserveronerrormarkdown", defaults_defaultserveronerrormarkdown},
		{"defaults_defaultserveronerrorsuddendeath", defaults_defaultserveronerrorsuddendeath},
		{"defaults_defaultserveronmarkeddownshutdow", defaults_defaultserveronmarkeddownshutdow},
		{"defaults_defaultserveronmarkedupshutdownb", defaults_defaultserveronmarkedupshutdownb},
		{"defaults_defaultserverpoollowconn384", defaults_defaultserverpoollowconn384},
		{"defaults_defaultserverpoolmaxconn0", defaults_defaultserverpoolmaxconn0},
		{"defaults_defaultserverpoolmaxconn1", defaults_defaultserverpoolmaxconn1},
		{"defaults_defaultserverpoolmaxconn100", defaults_defaultserverpoolmaxconn100},
		{"defaults_defaultserverpoolpurgedelay0", defaults_defaultserverpoolpurgedelay0},
		{"defaults_defaultserverpoolpurgedelay5", defaults_defaultserverpoolpurgedelay5},
		{"defaults_defaultserverpoolpurgedelay500", defaults_defaultserverpoolpurgedelay500},
		{"defaults_defaultserverport27015", defaults_defaultserverport27015},
		{"defaults_defaultserverport27016", defaults_defaultserverport27016},
		{"defaults_defaultserverprotoh2", defaults_defaultserverprotoh2},
		{"defaults_defaultserverproxyv2optionsssl", defaults_defaultserverproxyv2optionsssl},
		{"defaults_defaultserverproxyv2optionssslce", defaults_defaultserverproxyv2optionssslce},
		{"defaults_defaultserverproxyv2optionssslce_", defaults_defaultserverproxyv2optionssslce_},
		{"defaults_defaultserverredirhttpimage1mydo", defaults_defaultserverredirhttpimage1mydo},
		{"defaults_defaultserverredirhttpsimage1myd", defaults_defaultserverredirhttpsimage1myd},
		{"defaults_defaultserverresolvenet100008", defaults_defaultserverresolvenet100008},
		{"defaults_defaultserverresolvenet100008100", defaults_defaultserverresolvenet100008100},
		{"defaults_defaultserverresolveoptsallowdup", defaults_defaultserverresolveoptsallowdup},
		{"defaults_defaultserverresolveoptsallowdup_", defaults_defaultserverresolveoptsallowdup_},
		{"defaults_defaultserverresolveoptsignorewe", defaults_defaultserverresolveoptsignorewe},
		{"defaults_defaultserverresolveoptspreventd", defaults_defaultserverresolveoptspreventd},
		{"defaults_defaultserverresolvepreferipv4", defaults_defaultserverresolvepreferipv4},
		{"defaults_defaultserverresolvepreferipv6", defaults_defaultserverresolvepreferipv6},
		{"defaults_defaultserverresolversmydns", defaults_defaultserverresolversmydns},
		{"defaults_defaultserverrise2", defaults_defaultserverrise2},
		{"defaults_defaultserverrise200", defaults_defaultserverrise200},
		{"defaults_defaultserversendproxy", defaults_defaultserversendproxy},
		{"defaults_defaultserversendproxyv2", defaults_defaultserversendproxyv2},
		{"defaults_defaultserversendproxyv2ssl", defaults_defaultserversendproxyv2ssl},
		{"defaults_defaultserversendproxyv2sslcn", defaults_defaultserversendproxyv2sslcn},
		{"defaults_defaultserverslowstart2000ms", defaults_defaultserverslowstart2000ms},
		{"defaults_defaultserversniTODO", defaults_defaultserversniTODO},
		{"defaults_defaultserversocks412700181", defaults_defaultserversocks412700181},
		{"defaults_defaultserversourceTODO", defaults_defaultserversourceTODO},
		{"defaults_defaultserverssl", defaults_defaultserverssl},
		{"defaults_defaultserversslmaxverSSLv3", defaults_defaultserversslmaxverSSLv3},
		{"defaults_defaultserversslmaxverTLSv10", defaults_defaultserversslmaxverTLSv10},
		{"defaults_defaultserversslmaxverTLSv11", defaults_defaultserversslmaxverTLSv11},
		{"defaults_defaultserversslmaxverTLSv12", defaults_defaultserversslmaxverTLSv12},
		{"defaults_defaultserversslmaxverTLSv13", defaults_defaultserversslmaxverTLSv13},
		{"defaults_defaultserversslminverSSLv3", defaults_defaultserversslminverSSLv3},
		{"defaults_defaultserversslminverTLSv10", defaults_defaultserversslminverTLSv10},
		{"defaults_defaultserversslminverTLSv11", defaults_defaultserversslminverTLSv11},
		{"defaults_defaultserversslminverTLSv12", defaults_defaultserversslminverTLSv12},
		{"defaults_defaultserversslminverTLSv13", defaults_defaultserversslminverTLSv13},
		{"defaults_defaultserversslreuse", defaults_defaultserversslreuse},
		{"defaults_defaultserverstick", defaults_defaultserverstick},
		{"defaults_defaultservertcput20ms", defaults_defaultservertcput20ms},
		{"defaults_defaultservertfo", defaults_defaultservertfo},
		{"defaults_defaultservertlstickets", defaults_defaultservertlstickets},
		{"defaults_defaultservertrackTODO", defaults_defaultservertrackTODO},
		{"defaults_defaultserververifyhostsitecom", defaults_defaultserververifyhostsitecom},
		{"defaults_defaultserververifynone", defaults_defaultserververifynone},
		{"defaults_defaultserververifyrequired", defaults_defaultserververifyrequired},
		{"defaults_defaultserverweight1", defaults_defaultserverweight1},
		{"defaults_defaultserverweight128", defaults_defaultserverweight128},
		{"defaults_defaultserverweight256", defaults_defaultserverweight256},
		{"defaults_defaultserverwsauto", defaults_defaultserverwsauto},
		{"defaults_defaultserverwsh1", defaults_defaultserverwsh1},
		{"defaults_defaultserverwsh2", defaults_defaultserverwsh2},
		{"defaults_emailalertfromadminexamplecom", defaults_emailalertfromadminexamplecom},
		{"defaults_emailalertlevelwarning", defaults_emailalertlevelwarning},
		{"defaults_emailalertmailerslocalmailers", defaults_emailalertmailerslocalmailers},
		{"defaults_emailalertmyhostnamesrv01example", defaults_emailalertmyhostnamesrv01example},
		{"defaults_emailalerttoabcd", defaults_emailalerttoabcd},
		{"defaults_emailalerttoazxy", defaults_emailalerttoazxy},
		{"defaults_emailalerttosupportexamplecom", defaults_emailalerttosupportexamplecom},
		{"defaults_errorfile400etchaproxyerrorfiles", defaults_errorfile400etchaproxyerrorfiles},
		{"defaults_errorfile403etchaproxyerrorfiles", defaults_errorfile403etchaproxyerrorfiles},
		{"defaults_errorfile408devnullworkaroundChr", defaults_errorfile408devnullworkaroundChr},
		{"defaults_errorfile503etchaproxyerrorfiles", defaults_errorfile503etchaproxyerrorfiles},
		{"defaults_errorfileserrorssection", defaults_errorfileserrorssection},
		{"defaults_errorfileserrorssection400", defaults_errorfileserrorssection400},
		{"defaults_errorfileserrorssection400401500", defaults_errorfileserrorssection400401500},
		{"defaults_errorloc302400httpwwwmyawesomesi", defaults_errorloc302400httpwwwmyawesomesi},
		{"defaults_errorloc302404httpwwwmyawesomesi", defaults_errorloc302404httpwwwmyawesomesi},
		{"defaults_errorloc302501errorpage", defaults_errorloc302501errorpage},
		{"defaults_errorloc303400httpwwwmyawesomesi", defaults_errorloc303400httpwwwmyawesomesi},
		{"defaults_errorloc303404httpwwwmyawesomesi", defaults_errorloc303404httpwwwmyawesomesi},
		{"defaults_errorloc303501errorpage", defaults_errorloc303501errorpage},
		{"defaults_hashtypeavalanche", defaults_hashtypeavalanche},
		{"defaults_hashtypeconsistent", defaults_hashtypeconsistent},
		{"defaults_hashtypeconsistentavalanche", defaults_hashtypeconsistentavalanche},
		{"defaults_hashtypeconsistentcrc32", defaults_hashtypeconsistentcrc32},
		{"defaults_hashtypeconsistentcrc32avalanche", defaults_hashtypeconsistentcrc32avalanche},
		{"defaults_hashtypeconsistentdjb2", defaults_hashtypeconsistentdjb2},
		{"defaults_hashtypeconsistentdjb2avalanche", defaults_hashtypeconsistentdjb2avalanche},
		{"defaults_hashtypeconsistentnone", defaults_hashtypeconsistentnone},
		{"defaults_hashtypeconsistentsdbm", defaults_hashtypeconsistentsdbm},
		{"defaults_hashtypeconsistentsdbmavalanche", defaults_hashtypeconsistentsdbmavalanche},
		{"defaults_hashtypeconsistentwt6", defaults_hashtypeconsistentwt6},
		{"defaults_hashtypeconsistentwt6avalanche", defaults_hashtypeconsistentwt6avalanche},
		{"defaults_hashtypemapbased", defaults_hashtypemapbased},
		{"defaults_hashtypemapbasedavalanche", defaults_hashtypemapbasedavalanche},
		{"defaults_hashtypemapbasedcrc32", defaults_hashtypemapbasedcrc32},
		{"defaults_hashtypemapbasedcrc32avalanche", defaults_hashtypemapbasedcrc32avalanche},
		{"defaults_hashtypemapbaseddjb2", defaults_hashtypemapbaseddjb2},
		{"defaults_hashtypemapbaseddjb2avalanche", defaults_hashtypemapbaseddjb2avalanche},
		{"defaults_hashtypemapbasedsdbm", defaults_hashtypemapbasedsdbm},
		{"defaults_hashtypemapbasedsdbmavalanche", defaults_hashtypemapbasedsdbmavalanche},
		{"defaults_hashtypemapbasedwt6", defaults_hashtypemapbasedwt6},
		{"defaults_hashtypemapbasedwt6avalanche", defaults_hashtypemapbasedwt6avalanche},
		{"defaults_httpcheckcommenttestcomment", defaults_httpcheckcommenttestcomment},
		{"defaults_httpcheckconnect", defaults_httpcheckconnect},
		{"defaults_httpcheckconnectaddr8888", defaults_httpcheckconnectaddr8888},
		{"defaults_httpcheckconnectalpnh2http11", defaults_httpcheckconnectalpnh2http11},
		{"defaults_httpcheckconnectcommenttestcomme", defaults_httpcheckconnectcommenttestcomme},
		{"defaults_httpcheckconnectdefault", defaults_httpcheckconnectdefault},
		{"defaults_httpcheckconnectlinger", defaults_httpcheckconnectlinger},
		{"defaults_httpcheckconnectport443addr8888s", defaults_httpcheckconnectport443addr8888s},
		{"defaults_httpcheckconnectport8080", defaults_httpcheckconnectport8080},
		{"defaults_httpcheckconnectprotoh2", defaults_httpcheckconnectprotoh2},
		{"defaults_httpcheckconnectsendproxy", defaults_httpcheckconnectsendproxy},
		{"defaults_httpcheckconnectsnihaproxy1wteu", defaults_httpcheckconnectsnihaproxy1wteu},
		{"defaults_httpcheckconnectssl", defaults_httpcheckconnectssl},
		{"defaults_httpcheckconnectviasocks4", defaults_httpcheckconnectviasocks4},
		{"defaults_httpcheckdisableon404", defaults_httpcheckdisableon404},
		{"defaults_httpcheckexpectcommenttestcommen", defaults_httpcheckexpectcommenttestcommen},
		{"defaults_httpcheckexpecterrorstatusL7RSPs", defaults_httpcheckexpecterrorstatusL7RSPs},
		{"defaults_httpcheckexpectminrecv50status20", defaults_httpcheckexpectminrecv50status20},
		{"defaults_httpcheckexpectokstatusL7OKstatu", defaults_httpcheckexpectokstatusL7OKstatu},
		{"defaults_httpcheckexpectonerrormylogforma", defaults_httpcheckexpectonerrormylogforma},
		{"defaults_httpcheckexpectonsuccessmylogfor", defaults_httpcheckexpectonsuccessmylogfor},
		{"defaults_httpcheckexpectrstatus5", defaults_httpcheckexpectrstatus5},
		{"defaults_httpcheckexpectrstringtag09afhtm", defaults_httpcheckexpectrstringtag09afhtm},
		{"defaults_httpcheckexpectstatus200", defaults_httpcheckexpectstatus200},
		{"defaults_httpcheckexpectstatuscode500stat", defaults_httpcheckexpectstatuscode500stat},
		{"defaults_httpcheckexpectstringSQLError", defaults_httpcheckexpectstringSQLError},
		{"defaults_httpcheckexpecttoutstatusL7TOUTs", defaults_httpcheckexpecttoutstatusL7TOUTs},
		{"defaults_httpchecksendcommenttestcomment", defaults_httpchecksendcommenttestcomment},
		{"defaults_httpchecksendmethGET", defaults_httpchecksendmethGET},
		{"defaults_httpchecksendmethGETurihealthver", defaults_httpchecksendmethGETurihealthver},
		{"defaults_httpchecksendstate", defaults_httpchecksendstate},
		{"defaults_httpchecksendurihealth", defaults_httpchecksendurihealth},
		{"defaults_httpchecksendurilfmylogformatbod", defaults_httpchecksendurilfmylogformatbod},
		{"defaults_httpchecksendverHTTP11", defaults_httpchecksendverHTTP11},
		{"defaults_httpchecksetvarcheckportint1234", defaults_httpchecksetvarcheckportint1234},
		{"defaults_httpchecksetvarfmtcheckportint12", defaults_httpchecksetvarfmtcheckportint12},
		{"defaults_httpcheckunsetvartxnfrom", defaults_httpcheckunsetvartxnfrom},
		{"defaults_httperrorstatus200contenttypetex", defaults_httperrorstatus200contenttypetex},
		{"defaults_httperrorstatus400", defaults_httperrorstatus400},
		{"defaults_httperrorstatus400contenttypetex", defaults_httperrorstatus400contenttypetex},
		{"defaults_httperrorstatus400contenttypetex_", defaults_httperrorstatus400contenttypetex_},
		{"defaults_httperrorstatus400contenttypetex__", defaults_httperrorstatus400contenttypetex__},
		{"defaults_httperrorstatus400contenttypetex___", defaults_httperrorstatus400contenttypetex___},
		{"defaults_httperrorstatus400contenttypetex____", defaults_httperrorstatus400contenttypetex____},
		{"defaults_httperrorstatus400defaulterrorfi", defaults_httperrorstatus400defaulterrorfi},
		{"defaults_httperrorstatus400errorfilemyfan", defaults_httperrorstatus400errorfilemyfan},
		{"defaults_httperrorstatus400errorfilesmyer", defaults_httperrorstatus400errorfilesmyer},
		{"defaults_httpreuseaggressive", defaults_httpreuseaggressive},
		{"defaults_httpreusealways", defaults_httpreusealways},
		{"defaults_httpreusenever", defaults_httpreusenever},
		{"defaults_httpreusesafe", defaults_httpreusesafe},
		{"defaults_httpsendnameheader", defaults_httpsendnameheader},
		{"defaults_httpsendnameheaderXMyAwesomeHead", defaults_httpsendnameheaderXMyAwesomeHead},
		{"defaults_loadserverstatefromfileglobal", defaults_loadserverstatefromfileglobal},
		{"defaults_loadserverstatefromfilelocal", defaults_loadserverstatefromfilelocal},
		{"defaults_loadserverstatefromfilenone", defaults_loadserverstatefromfilenone},
		{"defaults_log1270011515formatrfc5424sample", defaults_log1270011515formatrfc5424sample},
		{"defaults_log1270011515formatrfc5424sample_", defaults_log1270011515formatrfc5424sample_},
		{"defaults_log1270011515len8192formatrfc542", defaults_log1270011515len8192formatrfc542},
		{"defaults_log1270011515len8192formatrfc542_", defaults_log1270011515len8192formatrfc542_},
		{"defaults_log1270011515sample12local0", defaults_log1270011515sample12local0},
		{"defaults_log1270011515sample16local2", defaults_log1270011515sample16local2},
		{"defaults_log127001514local0noticenoticesa", defaults_log127001514local0noticenoticesa},
		{"defaults_log127001514local0noticeonlysend", defaults_log127001514local0noticeonlysend},
		{"defaults_logglobal", defaults_logglobal},
		{"defaults_logstderrformatrawdaemonnoticese", defaults_logstderrformatrawdaemonnoticese},
		{"defaults_logstdoutformatrawdaemonsendever", defaults_logstdoutformatrawdaemonsendever},
		{"defaults_logstdoutformatshortdaemonsendlo", defaults_logstdoutformatshortdaemonsendlo},
		{"defaults_monitorurihaproxytest", defaults_monitorurihaproxytest},
		{"defaults_nolog", defaults_nolog},
		{"defaults_optionhttpchkOPTIONSHTTP11rnHost", defaults_optionhttpchkOPTIONSHTTP11rnHost},
		{"defaults_optionhttpchkmethoduri", defaults_optionhttpchkmethoduri},
		{"defaults_optionhttpchkmethoduriversion", defaults_optionhttpchkmethoduriversion},
		{"defaults_optionhttpchkuri", defaults_optionhttpchkuri},
		{"defaults_optionhttprestrictreqhdrnamesdel", defaults_optionhttprestrictreqhdrnamesdel},
		{"defaults_optionhttprestrictreqhdrnamespre", defaults_optionhttprestrictreqhdrnamespre},
		{"defaults_optionhttprestrictreqhdrnamesrej", defaults_optionhttprestrictreqhdrnamesrej},
		{"defaults_optionoriginalto", defaults_optionoriginalto},
		{"defaults_optionoriginaltocomment", defaults_optionoriginaltocomment},
		{"defaults_optionoriginaltoexcept127001", defaults_optionoriginaltoexcept127001},
		{"defaults_optionoriginaltoexcept127001comm", defaults_optionoriginaltoexcept127001comm},
		{"defaults_optionoriginaltoexcept127001head", defaults_optionoriginaltoexcept127001head},
		{"defaults_optionoriginaltoheaderXClientDst", defaults_optionoriginaltoheaderXClientDst},
		{"defaults_persistrdpcookie", defaults_persistrdpcookie},
		{"defaults_persistrdpcookiecookies", defaults_persistrdpcookiecookies},
		{"defaults_source0000usesrcclientip", defaults_source0000usesrcclientip},
		{"defaults_source0000usesrchdripxforwardedf", defaults_source0000usesrchdripxforwardedf},
		{"defaults_source1921681200", defaults_source1921681200},
		{"defaults_source192168120080usesrcclient", defaults_source192168120080usesrcclient},
		{"defaults_source192168120080usesrcclientip", defaults_source192168120080usesrcclientip},
		{"defaults_source1921681200interfacename", defaults_source1921681200interfacename},
		{"defaults_source1921681200usesrc1921681201", defaults_source1921681200usesrc1921681201},
		{"defaults_source1921681200usesrcclient", defaults_source1921681200usesrcclient},
		{"defaults_source1921681200usesrcclientip", defaults_source1921681200usesrcclientip},
		{"defaults_source1921681200usesrchdriphdr", defaults_source1921681200usesrchdriphdr},
		{"defaults_source1921681200usesrchdriphdroc", defaults_source1921681200usesrchdriphdroc},
		{"defaults_statsauthadmin1AdMiN123", defaults_statsauthadmin1AdMiN123},
		{"defaults_statsbindprocess1234", defaults_statsbindprocess1234},
		{"defaults_statsbindprocess14", defaults_statsbindprocess14},
		{"defaults_statsbindprocessall", defaults_statsbindprocessall},
		{"defaults_statsbindprocesseven", defaults_statsbindprocesseven},
		{"defaults_statsbindprocessodd", defaults_statsbindprocessodd},
		{"defaults_statsenable", defaults_statsenable},
		{"defaults_statshideversion", defaults_statshideversion},
		{"defaults_statsmaxconn10", defaults_statsmaxconn10},
		{"defaults_statsrealmHAProxyStatistics", defaults_statsrealmHAProxyStatistics},
		{"defaults_statsrefresh10s", defaults_statsrefresh10s},
		{"defaults_statsscope", defaults_statsscope},
		{"defaults_statsshowdescMasternodeforEurope", defaults_statsshowdescMasternodeforEurope},
		{"defaults_statsshowlegends", defaults_statsshowlegends},
		{"defaults_statsshowmodules", defaults_statsshowmodules},
		{"defaults_statsshownode", defaults_statsshownode},
		{"defaults_statsshownodeEurope1", defaults_statsshownodeEurope1},
		{"defaults_statsuriadminstats", defaults_statsuriadminstats},
		{"defaults_tcpcheckcommenttestcomment", defaults_tcpcheckcommenttestcomment},
		{"defaults_tcpcheckconnect", defaults_tcpcheckconnect},
		{"defaults_tcpcheckconnectport110linger", defaults_tcpcheckconnectport110linger},
		{"defaults_tcpcheckconnectport143", defaults_tcpcheckconnectport143},
		{"defaults_tcpcheckconnectport443ssl", defaults_tcpcheckconnectport443ssl},
		{"defaults_tcpcheckexpectstringOK", defaults_tcpcheckexpectstringOK},
		{"defaults_tcpcheckexpectstringOKIMAP4ready", defaults_tcpcheckexpectstringOKIMAP4ready},
		{"defaults_tcpcheckexpectstringOKPOP3ready", defaults_tcpcheckexpectstringOKPOP3ready},
		{"defaults_tcpcheckexpectstringPONG", defaults_tcpcheckexpectstringPONG},
		{"defaults_tcpcheckexpectstringrolemaster", defaults_tcpcheckexpectstringrolemaster},
		{"defaults_tcpchecksendPINGrn", defaults_tcpchecksendPINGrn},
		{"defaults_tcpchecksendPINGrncommenttestcom", defaults_tcpchecksendPINGrncommenttestcom},
		{"defaults_tcpchecksendQUITrn", defaults_tcpchecksendQUITrn},
		{"defaults_tcpchecksendQUITrncommenttestcom", defaults_tcpchecksendQUITrncommenttestcom},
		{"defaults_tcpchecksendbinarylftesthexfmt", defaults_tcpchecksendbinarylftesthexfmt},
		{"defaults_tcpchecksendbinarylftesthexfmtco", defaults_tcpchecksendbinarylftesthexfmtco},
		{"defaults_tcpchecksendbinarytesthexstring", defaults_tcpchecksendbinarytesthexstring},
		{"defaults_tcpchecksendbinarytesthexstringc", defaults_tcpchecksendbinarytesthexstringc},
		{"defaults_tcpchecksendinforeplicationrn", defaults_tcpchecksendinforeplicationrn},
		{"defaults_tcpchecksendlftestfmt", defaults_tcpchecksendlftestfmt},
		{"defaults_tcpchecksendlftestfmtcommenttest", defaults_tcpchecksendlftestfmtcommenttest},
		{"defaults_tcpchecksetvarcheckportint1234", defaults_tcpchecksetvarcheckportint1234},
		{"defaults_tcpchecksetvarfmtchecknameH", defaults_tcpchecksetvarfmtchecknameH},
		{"defaults_tcpchecksetvarfmttxnfromaddrsrcs", defaults_tcpchecksetvarfmttxnfromaddrsrcs},
		{"defaults_tcpcheckunsetvartxnfrom", defaults_tcpcheckunsetvartxnfrom},
		{"defaults_uniqueidformatXocicpfifpTsrtpid", defaults_uniqueidformatXocicpfifpTsrtpid},
		{"defaults_uniqueidformatXocpfifpTsrtpid", defaults_uniqueidformatXocpfifpTsrtpid},
		{"defaults_uniqueidformatXofifpTsrtpid", defaults_uniqueidformatXofifpTsrtpid},
		{"defaults_uniqueidheaderXUniqueID", defaults_uniqueidheaderXUniqueID},
	}
	for _, config := range tests {
		t.Run(config.Name, func(t *testing.T) {
			t.Parallel()
			var buffer bytes.Buffer
			buffer.WriteString(config.Config)
			p, err := parser.New(options.Reader(&buffer))
			if err != nil {
				t.Fatalf(err.Error())
			}
			result := p.String()
			if result != config.Config {
				compare(t, config.Config, result)
				t.Error("======== ORIGINAL =========")
				t.Error(config.Config)
				t.Error("======== RESULT ===========")
				t.Error(result)
				t.Error("===========================")
				t.Fatalf("configurations does not match")
			}
		})
	}
}
