package oracle

import (
	"context"
	"encoding/xml"
	"os"
	"testing"

	"github.com/quay/claircore/test/log"

	"github.com/quay/goval-parser/oval"
)

func TestParse(t *testing.T) {
	t.Parallel()
	l := log.TestLogger(t)
	ctx := l.WithContext(context.Background())
	u, err := NewUpdater(-1)
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Open("testdata/com.oracle.elsa-2018.xml")
	if err != nil {
		t.Fatal(err)
	}

	vs, err := u.ParseContext(ctx, f)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("found %d vulnerabilities", len(vs))
	if got, want := len(vs), 3128; got != want {
		t.Fatalf("got: %d vulnerabilities, want: %d vulnerabilities", got, want)
	}
}

var ovalDef = oval.Definition{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "definition"},
	ID:    "oval:com.oracle.elsa:def:20162594",
	Class: "patch",
	Title: "\nELSA-2016-2594:  389-ds-base security, bug fix, and enhancement update (MODERATE)\n",
	Affecteds: []oval.Affected{
		{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "affected"},
			Family:    "unix",
			Platforms: []string{"Oracle Linux 7"}}},
	References: []oval.Reference{
		{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "reference"},
			Source: "elsa",
			RefID:  "ELSA-2016-2594",
			RefURL: "http://linux.oracle.com/errata/ELSA-2016-2594.html"},
		{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "reference"},
			Source: "CVE",
			RefID:  "CVE-2016-4992",
			RefURL: "http://linux.oracle.com/cve/CVE-2016-4992.html"},
		{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "reference"},
			Source: "CVE",
			RefID:  "CVE-2016-5405",
			RefURL: "http://linux.oracle.com/cve/CVE-2016-5405.html"},
		{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "reference"},
			Source: "CVE",
			RefID:  "CVE-2016-5416",
			RefURL: "http://linux.oracle.com/cve/CVE-2016-5416.html"}},
	Description: "\n[1.3.5.10-11]\n- Release 1.3.5.10-11\n- Resolves: bug 1321124 - Replication changelog can incorrectly skip over updates\n\n[1.3.5.10-10]\n- Release 1.3.5.10-10\n- Resolves: bug 1370300 - set proper update status to replication agreement in case of failure (DS 48957)\n- Resolves: bug 1209094 - Allow logging of rejected changes (DS 48969)\n\n[1.3.5.10-9]\n- Release 1.3.5.10-9\n- Resolves: bug 1364190 - Change example in /etc/sysconfig/dirsrv to use tcmalloc (DS 48950)\n- Resolves: bug 1366828 - audit on failure doesn't work if attribute nsslapd-auditlog-logging-enabled is NOT enabled (DS 48958)\n- Resolves: bug 1368520 - Crash in import_wait_for_space_in_fifo() (DS 48960)\n- Resolves: bug 1368956 - man page of ns-accountstatus.pl shows redundant entries for -p port option\n- Resolves: bug 1369537 - passwordMinAge attribute doesn't limit the minimum age of the password (DS 48967)\n- Resolves: bug 1369570 - cleanallruv changelog cleaning incorrectly impacts all backend (DS 48964)\n- Resolves: bug 1369425 - ACI behaves erratically (DS 48972)\n- Resolves: bug 1370300 - set proper update status to replication agreement in case of failure (DS 48957)\n- Resolves: bug 1209094 - Allow logging of rejected changes (DS 48969)\n- Resolves: bug 1371283 - Server Side Sorting crashes the server. (DS 48970)\n- Resolves: bug 1371284 - Disabling CLEAR password storage scheme will crash server when setting a password (DS 48975)\n\n[1.3.5.10-8]\n- Release 1.3.5.10-8\n- Resolves: bug 1321124 - Replication changelog can incorrectly skip over updates (DS 48954)\n- Resolves: bug 1364190 - Change example in /etc/sysconfig/dirsrv to use tcmalloc (DS 48950)\n- Resolves: bug 1366561 - ns-accountstatus.pl giving error even 'No such object (32)' (DS 48956)\n\n[1.3.5.10-7]\n- Release 1.3.5.10-7\n- Resolves: bug 1316580 - dirsrv service doesn't ask for pin when pin.txt is missing (DS 48450)\n- Resolves: bug 1360976 - fixing a compiler warning\n\n[1.3.5.10-6]\n- Release 1.3.5.10-6\n- Resolves: bug 1326077 - Page result search should return empty cookie if there is no returned entry (DS 48928)\n- Resolves: bug 1360447 - nsslapd-workingdir is empty when ns-slapd is started by systemd (DS 48939)\n- Resolves: bug 1360327 - remove-ds.pl deletes an instance even if wrong prefix was specified (DS 48934)\n- Resolves: bug 1349815 - DS logs have warning:ancestorid not indexed for all CS subsystems (DS 48940)\n- Resolves: bug 1329061 - 389-ds-base-1.3.4.0-29.el7_2 'hang' (DS 48882)\n- Resolves: bug 1360976 - EMBARGOED CVE-2016-5405 389-ds-base: Password verification vulnerable to timing attack\n- Resolves: bug 1361134 - When fine-grained policy is applied, a sub-tree has a priority over a user while changing password (DS 48943)\n- Resolves: bug 1361321 - Duplicate collation entries (DS 48936)\n- Resolves: bug 1316580 - dirsrv service doesn't ask for pin when pin.txt is missing (DS 48450)\n- Resolves: bug 1350799 - CVE-2016-4992 389-ds-base: Information disclosure via repeat\n\n[1.3.5.10-5]\n- Release 1.3.5.10-5\n- Resolves: bug 1333184 - (389-ds-base-1.3.5) Fixing coverity issues. (DS 48919)\n\n[1.3.5.10-4]\n- Release 1.3.5.10-4\n- Resolves: bug 1209128 - [RFE] Add a utility to get the status of Directory Server instances (DS 48144)\n- Resolves: bug 1333184 - (389-ds-base-1.3.5) Fixing coverity issues. (DS 48919)\n- Resolves: bug 1350799 - CVE-2016-4992 389-ds-base: Information disclosure via repeat\n- Resolves: bug 1354660 - flow control in replication also blocks receiving results (DS 48767)\n- Resolves: bug 1356261 - Fixup tombstone task needs to set proper flag when updating (DS 48924)\n- Resolves: bug 1355760 - ns-slapd crashes during the deletion of backend (DS 48922)\n- Resolves: bug 1353629 - DS shuts down automatically if dnaThreshold is set to 0 in a MMR setup (DS 48916)\n- Resolves: bug 1355879 - nunc-stans: ns-slapd crashes during startup with SIGILL on AMD Opteron 280 (DS 48925)\n\n[1.3.5.10-3]\n- Release 1.3.5.10-3\n- Resolves: bug 1354374 - Fixing the tarball version in the sources file.\n\n[1.3.5.10-2]\n- Release 1.3.5.10-2\n- Resolves: bug 1353714 - If a cipher is disabled do not attempt to look it up (DS 48743)\n- Resolves: bug 1353592 - Setup-ds.pl --update fails - regression (DS 48755)\n- Resolves: bug 1353544 - db2bak.pl task enters infinitive loop when bak fs is almost full (DS 48914)\n- Resolves: bug 1354374 - Upgrade to 389-ds-base >= 1.3.5.5 doesn't install 389-ds-base-snmp (DS 48918)\n\n[1.3.5.10-1]\n- Release 1.3.5.10-1\n- Resolves: bug 1333184 - (389-ds-base-1.3.5) Fixing coverity issues. (DS 48905)\n\n[1.3.5.9-1]\n- Release 1.3.5.9-1\n- Resolves: bug 1349571 - Improve MMR replication convergence (DS 48636)\n- Resolves: bug 1304682 - 'stale' automember rule (associated to a removed group) causes discrepancies in the database (DS 48637)\n- Resolves: bug 1314956 - moving an entry cause next on-line init to skip entry has no parent, ending at line 0 of file '(bulk import)' (DS 48755)\n- Resolves: bug 1316731 - syncrepl search returning error 329; plugin sending a bad error code (DS 48904)\n- Resolves: bug 1346741 - ns-slapd crashes during the shutdown after adding attribute with a matching rule  (DS 48891)\n- Resolves: bug 1349577 - Values of dbcachetries/dbcachehits in cn=monitor could overflow. (DS 48899)\n- Resolves: bug 1272682 - nunc-stans: ns-slapd killed by SIGABRT (DS 48898)\n- Resolves: bug 1346043 - repl-monitor displays colors incorrectly for the time lag > 60 min (DS 47538)\n- Resolves: bug 1350632 - ns-slapd shutdown crashes if pwdstorageschema name is from stack. (DS 48902)\n\n[1.3.5.8-1]\n- Release 1.3.5.8-1\n- Resolves: bug 1290101 - proxyauth support does not work when bound as directory  manager (DS 48366)\n\n[1.3.5.7-1]\n- Release 1.3.5.7-1\n- Resolves: bug 1196282 - substring index with nssubstrbegin: 1 is not being used with filters like (attr=x*) (DS 48109)\n- Resolves: bug 1303794 - Import readNSState.py from RichM's repo (DS 48449)\n- Resolves: bug 1290101 - proxyauth support does not work when bound as directory  manager (DS 48366)\n- Resolves: bug 1338872 - Wrong result code display in audit-failure log (DS 48892)\n- Resolves: bug 1346043 - repl-monitor displays colors incorrectly for the time lag > 60 min (DS 47538)\n- Resolves: bug 1346741 - ns-slapd crashes during the shutdown after adding attribute with a matching rule  (DS 48891)\n- Resolves: bug 1347407 - By default aci can be read by anonymous (DS 48354)\n- Resolves: bug 1347412 - cn=SNMP,cn=config entry can be read by anonymous (DS 48893)\n\n[1.3.5.6-1]\n- Release 1.3.5.6-1\n- Resolves: bug 1273549 - [RFE] Improve timestamp resolution in logs (DS 47982)\n- Resolves: bug 1321124 - Replication changelog can incorrectly skip over updates (DS 48766, DS 48636)\n- Resolves: bug 1233926 - 'matching rules' in ACI's 'bind rules not fully evaluated (DS 48234)\n- Resolves: bug 1346165 - 389-ds-base-1.3.5.5-1.el7.x86_64 requires policycoreutils-py\n\n[1.3.5.5-1]\n- Release 1.3.5.5-1\n- Resolves: bug 1018944 - [RFE] Enhance password change tracking (DS 48833)\n- Resolves: bug 1344414 - [RFE] adding pre/post extop ability (DS 48880)\n- Resolves: bug 1303794 - Import readNSState.py from RichM's repo (DS 48449)\n- Resolves: bug 1257568 - /usr/lib64/dirsrv/libnunc-stans.so is owned by both -libs and -devel (DS 48404)\n- Resolves: bug 1314956 - moving an entry cause next on-line init to skip entry has no parent, ending at line 0 of file '(bulk import)' (DS 48755)\n- Resolves: bug 1342609 - At startup DES to AES password conversion causes timeout in start script (DS 48862)\n- Resolves: bug 1316328 - search returns no entry when OR filter component contains non readable attribute (DS 48275)\n- Resolves: bug 1280456 - setup-ds should detect if port is already defined (DS 48336)\n- Resolves: bug 1312557 - dirsrv service fails to start when nsslapd-listenhost is configured (DS 48747)\n- Resolves: bug 1326077 - Page result search should return empty cookie if there is no returned entry (DS 48752)\n- Resolves: bug 1340307 - Running db2index with no options breaks replication (DS 48854)\n- Resolves: bug 1337195 - Regression introduced in matching rules by DS 48746 (DS 48844)\n- Resolves: bug 1335492 - Modifier's name is not recorded in the audit log with modrdn and moddn operations (DS 48834)\n- Resolves: bug 1316741 - ldctl should support -H with ldap uris (DS 48754)\n\n[1.3.5.4-1]\n- release 1.3.5.4-1\n- Resolves: bug 1334455 - db2ldif is not taking into account multiple suffixes or backends (DS 48828)\n- Resolves: bug 1241563 - The 'repl-monitor' web page does not display 'year' in date. (DS 48220)\n- Resolves: bug 1335618 - Server ram sanity checks work in isolation (DS 48617)\n- Resolves: bug 1333184 - (389-ds-base-1.3.5) Fixing coverity issues. (DS 48837)\n\n[1.3.5.3-1]\n- release 1.3.5.3-1\n- Resolves: bug 1209128 - [RFE] Add a utility to get the status of Directory Server instances (DS 48144)\n- Resolves: bug 1332533 - ns-accountstatus.pl gives error message on execution along with results. (DS 48815)\n- Resolves: bug 1332709 - password history is not updated when an admin resets the password (DS 48813)\n- Resolves: bug 1333184 - (389-ds-base-1.3.5) Fixing coverity issues. (DS 48822)\n- Resolves: bug 1333515 - Enable DS to offer weaker DH params in NSS  (DS 48798)\n\n[1.3.5.2-1]\n- release 1.3.5.2-1\n- Resolves: bug 1270020 - Rebase 389-ds-base to 1.3.5 in RHEL-7.3\n- Resolves: bug 1288229 - many attrlist_replace errors in connection with cleanallruv (DS 48283)\n- Resolves: bug 1315893 - License tag does not match actual license of code (DS 48757)\n- Resolves: bug 1320715 - DES to AES password conversion fails if a backend is empty (DS 48777)\n- Resolves: bug 190862  - [RFE] Default password syntax settings don't work with fine-grained policies (DS 142)\n- Resolves: bug 1018944 - [RFE] Enhance password change tracking (DS 548)\n- Resolves: bug 1143066 - The dirsrv user/group should be created in rpm %pre, and ideally with fixed uid/gid (DS 48285)\n- Resolves: bug 1153758 - [RFE] Support SASL/GSSAPI when ns-slapd is behind a load-balancer (DS 48332)\n- Resolves: bug 1160902 - search, matching rules and filter error 'unsupported type 0xA9' (DS 48016)\n- Resolves: bug 1186512 - High memory fragmentation observed in ns-slapd; OOM-Killer invoked (DS 48377, 48129)\n- Resolves: bug 1196282 - substring index with nssubstrbegin: 1 is not being used with filters like (attr=x*) (DS 48109)\n- Resolves: bug 1209094 - [RFE] Allow logging of rejected changes (DS 48145, 48280)\n- Resolves: bug 1209128 - [RFE] Add a utility to get the status of Directory Server instances (DS 48144)\n- Resolves: bug 1210842 - [RFE] Add PIDFile option to systemd service file (DS 47951)\n- Resolves: bug 1223510 - [RFE] it could be nice to have nsslapd-maxbersize default to bigger than 2Mb (DS 48326)\n- Resolves: bug 1229799 - ldclt-bin killed by SIGSEGV (DS 48289)\n- Resolves: bug 1249908 - No validation check for the value for nsslapd-db-locks. (DS 48244)\n- Resolves: bug 1254887 - No man page entry for - option '-u' of dbgen.pl for adding group entries with uniquemembers (DS 48290)\n- Resolves: bug 1255557 - db2index creates index entry from deleted records (DS 48252)\n- Resolves: bug 1258610 - total update request must not be lost (DS 48255)\n- Resolves: bug 1258611 - dna plugin needs to handle binddn groups for authorization (DS 48258)\n- Resolves: bug 1259624 - [RFE] Provide a utility to detect accounts locked due to inactivity (DS 48269)\n- Resolves: bug 1259950 - Add config setting to MemberOf Plugin to add required objectclass got memberOf attribute (DS 48267)\n- Resolves: bug 1266510 - Linked Attributes plug-in - wrong behaviour when adding valid and broken links (DS 48295)\n- Resolves: bug 1266532 - Linked Attributes plug-in - won't update links after MODRDN operation (DS 48294)\n- Resolves: bug 1267750 - pagedresults - when timed out, search results could have been already freed. (DS 48299)\n- Resolves: bug 1269378 - ds-logpipe.py with wrong arguments - python exception in the output (DS 48302)\n- Resolves: bug 1271330 - nunc-stans: Attempt to release connection that is not acquired (DS 48311)\n- Resolves: bug 1272677 - nunc stans: ns-slapd killed by SIGTERM\n- Resolves: bug 1272682 - nunc-stans: ns-slapd killed by SIGABRT\n- Resolves: bug 1273142 - crash in Managed Entry plugin (DS 48312)\n- Resolves: bug 1273549 - [RFE] Improve timestamp resolution in logs (DS 47982)\n- Resolves: bug 1273550 - Deadlock between two MODs on the same entry between entry cache and backend lock (DS 47978)\n- Resolves: bug 1273555 - deadlock in mep delete post op (DS 47976)\n- Resolves: bug 1273584 - lower password history minimum to 1 (DS 48394)\n- Resolves: bug 1275763 - [RFE] add setup-ds.pl option to disable instance specific scripts (DS 47840)\n- Resolves: bug 1276072 - [RFE] Allow RHDS to be setup using a DNS CNAME alias for General.FullMachineName (DS 48328)\n- Resolves: bug 1278567 - SimplePagedResults -- abandon could happen between the abandon check and sending results (DS 48338)\n- Resolves: bug 1278584 - Share nsslapd-threadnumber in the case nunc-stans is enabled, as well. (DS 48339)\n- Resolves: bug 1278755 - deadlock on connection mutex (DS 48341)\n- Resolves: bug 1278987 - Cannot upgrade a consumer to supplier in a multimaster environment (DS 48325)\n- Resolves: bug 1280123 - acl - regression - trailing ', (comma)' in macro matched value is not removed. (DS 48344)\n- Resolves: bug 1290111 - [RFE] Support for rfc3673 '+' to return operational attributes (DS 48363)\n- Resolves: bug 1290141 - With exhausted range, part of DNA shared configuration is deleted after server restart (DS 48362)\n- Resolves: bug 1290242 - SimplePagedResults -- in the search error case, simple paged results slot was not released. (DS 48375)\n- Resolves: bug 1290600 - The 'eq' index does not get updated properly when deleting and re-adding attributes in the same ldapmodify operation (DS 48370)\n- Resolves: bug 1295947 - 389-ds hanging after a few minutes of operation (DS 48406, revert 48338)\n- Resolves: bug 1296310 - ldclt - segmentation fault error while binding (DS 48400)\n- Resolves: bug 1299758 - CVE-2016-0741 389-ds-base: Worker threads do not detect abnormally closed connections causing DoS [rhel-7.3]\n- Resolves: bug 1301097 - logconv.pl displays negative operation speeds (DS 48446)\n- Resolves: bug 1302823 - Crash in slapi_get_object_extension (DS 48536)\n- Resolves: bug 1303641 - heap corruption at schema replication. (DS 48492)\n- Resolves: bug 1307151 - keep alive entries can break replication (DS 48445)\n- Resolves: bug 1310848 - Supplier can skip a failing update, although it should retry. (DS 47788)\n- Resolves: bug 1314557 - change severity of some messages related to 'keep alive' enties (DS 48420)\n- Resolves: bug 1316580 - dirsrv service doesn't ask for pin when pin.txt is missing (DS 48450)\n- Resolves: bug 1316742 - no plugin calls in tombstone purging (DS 48759)\n- Resolves: bug 1319329 - [RFE] add nsslapd-auditlog-logging-enabled: off to template-dse.ldif (DS 48145)\n- Resolves: bug 1320295 - If nsSSL3 is on, even if SSL v3 is not really enabled, a confusing message is logged. (DS 48775)\n- Resolves: bug 1326520 - db2index uses a buffer size derived from dbcachesize (DS 48383)\n- Resolves: bug 1328936 - objectclass values could be dropped on the consumer (DS 48799)\n- Resolves: bug 1287475 - [RFE] response control for password age should be sent by default by RHDS (DS 48369)\n- Resolves: bug 1331343 - Paged results search returns the blank list of entries (DS 48808)\n\n",
	Advisory: oval.Advisory{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "advisory"},
		Severity: "MODERATE",
		Cves: []oval.Cve{
			{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "cve"},
				CveID: "CVE-2016-4992",
				Href:  "http://linux.oracle.com/cve/CVE-2016-4992.html"},
			{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "cve"},
				CveID: "CVE-2016-5405",
				Href:  "http://linux.oracle.com/cve/CVE-2016-5405.html"},
			{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "cve"},
				CveID: "CVE-2016-5416",
				Href:  "http://linux.oracle.com/cve/CVE-2016-5416.html"},
		},
		Issued: struct {
			Date string "xml:\"date,attr\""
		}{Date: "2016-11-09"},
	},
	Criteria: oval.Criteria{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criteria"},
		Operator: "AND",
		Criterias: []oval.Criteria{
			{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criteria"},
				Operator: "OR",
				Criterias: []oval.Criteria{
					{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criteria"},
						Operator: "AND",
						Criterions: []oval.Criterion{
							{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criterion"},
								Negate:  false,
								TestRef: "oval:com.oracle.elsa:tst:20162594002",
								Comment: "389-ds-base is earlier than 0:1.3.5.10-11.el7"},
							{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criterion"},
								Negate:  false,
								TestRef: "oval:com.oracle.elsa:tst:20162594003",
								Comment: "389-ds-base is signed with the Oracle Linux 7 key"}}},
					{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criteria"},
						Operator: "AND",
						Criterions: []oval.Criterion{
							{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criterion"},
								Negate:  false,
								TestRef: "oval:com.oracle.elsa:tst:20162594004",
								Comment: "389-ds-base-libs is earlier than 0:1.3.5.10-11.el7"},
							{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5",
								Local: "criterion"},
								Negate:  false,
								TestRef: "oval:com.oracle.elsa:tst:20162594005",
								Comment: "389-ds-base-libs is signed with the Oracle Linux 7 key"}}},
					{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criteria"},
						Operator: "AND",
						Criterions: []oval.Criterion{
							{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criterion"},
								Negate:  false,
								TestRef: "oval:com.oracle.elsa:tst:20162594006",
								Comment: "389-ds-base-snmp is earlier than 0:1.3.5.10-11.el7"},
							{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criterion"},
								Negate:  false,
								TestRef: "oval:com.oracle.elsa:tst:20162594007",
								Comment: "389-ds-base-snmp is signed with the Oracle Linux 7 key"}}},
					{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criteria"},
						Operator: "AND",
						Criterions: []oval.Criterion{
							{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5",
								Local: "criterion"},
								Negate:  false,
								TestRef: "oval:com.oracle.elsa:tst:20162594008",
								Comment: "389-ds-base-devel is earlier than 0:1.3.5.10-11.el7"},
							{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criterion"},
								Negate:  false,
								TestRef: "oval:com.oracle.elsa:tst:20162594009",
								Comment: "389-ds-base-devel is signed with the Oracle Linux 7 key"}}}},
				Criterions: []oval.Criterion{
					{XMLName: xml.Name{Space: "http://oval.mitre.org/XMLSchema/oval-definitions-5", Local: "criterion"},
						Negate:  false,
						TestRef: "oval:com.oracle.elsa:tst:20162594001",
						Comment: "Oracle Linux 7 is installed"}}}}}}
