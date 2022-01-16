<script>
    import Home from "../admin/Home.svelte";
    import Entry from "../admin/Entry.svelte";
    let listAdmin = [];
    let record = "";
    let totalrecord = 0;
    let sData = "";
    let admin_username = "";
    let token = localStorage.getItem("token");
    let client_key = "";
    let akses_page = true;
    let admin_listrule = [];
    let admin_listip = [];
    let admin_name_field = "";
    let admin_type_field = "";
    let admin_idrule_field = "";
    let admin_status_field = "";
    let admin_create_field = "";
    let admin_update_field = "";

    const handleBackHalaman = () => {
        admin_username = "";
        sData = "";
        listAdmin = [];
        admin_listrule = [];
        admin_listip = [];
        handleRefreshData("all");
    };
    const handleRefreshData = (e) => {
        listAdmin = [];
        admin_listrule = [];
        totalrecord = 0;
        setTimeout(function () {
            initAdmin();
        }, 500);
    };
    const handleEditData = (e) => {
        admin_username = e.detail.e;
        admin_type_field = e.detail.f;
        sData = "Edit";
        editAdmin(admin_username);
    };
    async function initapp() {
        const res = await fetch("/api/home", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "ADMIN-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            // setTimeout(function(){ initPasaran() }, 1000);
            initAdmin();
        }
    }
    async function initAdmin() {
        const res = await fetch("/api/alladmin", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            let recordlistrule = json.listruleadmin;
            let css_status = "";
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    if (record[i]["admin_status"] == "ACTIVE") {
                        css_status =
                            "background:#8BC34A;color:black;font-weight:bold;";
                    } else {
                        css_status =
                            "background:#E91E63;font-size:12px;font-weight:bold;color:white;";
                    }

                    listAdmin = [
                        ...listAdmin,
                        {
                            admin_no: record[i]["admin_no"],
                            admin_username: record[i]["admin_username"],
                            admin_nama: record[i]["admin_nama"],
                            admin_tipe: record[i]["admin_tipe"],
                            admin_rule: record[i]["admin_rule"],
                            admin_timezone: record[i]["admin_timezone"],
                            admin_joindate: record[i]["admin_joindate"],
                            admin_lastlogin: record[i]["admin_lastlogin"],
                            admin_lastipaddres: record[i]["admin_lastipaddres"],
                            admin_status: record[i]["admin_status"],
                            admin_statuscss: css_status,
                        },
                    ];
                }
            }
            if (recordlistrule != null) {
                for (let i = 0; i < recordlistrule.length; i++) {
                    admin_listrule = [
                        ...admin_listrule,
                        {
                            adminrule_idruleadmin:
                                recordlistrule[i]["adminrule_idruleadmin"],
                            adminrule_name: recordlistrule[i]["adminrule_name"],
                        },
                    ];
                }
            }
        } else {
            logout();
        }
    }
    async function editAdmin(e) {
        clearEdit();
        const res = await fetch("/api/editadmin", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                username: e,
            }),
        });
        const json = await res.json();
        let record = json.record;
        let recordlistrule = json.listruleadmin;
        let recordlistip = json.listip;

        if (json.status === 400) {
            logout();
        } else {
            for (let i = 0; i < record.length; i++) {
                admin_name_field = record[i]["admin_nama"];
                admin_idrule_field = parseInt(record[i]["admin_idrule"]);
                admin_status_field = record[i]["admin_status"];
                admin_create_field = record[i]["admin_create"];
                admin_update_field = record[i]["admin_update"];
            }
            if (recordlistrule != null) {
                for (let i = 0; i < recordlistrule.length; i++) {
                    admin_listrule = [
                        ...admin_listrule,
                        {
                            adminrule_idruleadmin:recordlistrule[i]["adminrule_idruleadmin"],
                            adminrule_name: recordlistrule[i]["adminrule_name"],
                        },
                    ];
                }
            }
            if (recordlistip != null) {
                let no = 0;
                for (let i = 0; i < recordlistip.length; i++) {
                    no = no + 1;
                    admin_listip = [
                        ...admin_listip,
                        {
                            adminiplist_no: no,
                            adminiplist_idcompiplist:recordlistip[i]["adminiplist_idcompiplist"],
                            adminiplist_iplist:recordlistip[i]["adminiplist_iplist"],
                        },
                    ];
                }
            }
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshIplist = (e) => {
        let username = e.detail;
        admin_listip = [];
        editAdmin(username);
    };
    const handleDeleteIplist = (e) => {
        let idcompiplist = e.detail.e;
        let username = e.detail.admin_username;
        deleteIpList(idcompiplist, username);
    };
    async function deleteIpList(e, x) {
        const res = await fetch("/api/deleteadminiplist", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idcompiplist: parseInt(e),
                username: x,
                page:"ADMIN-SAVE",
            }),
        });
        const json = await res.json();

        if (json.status == 200) {
            admin_listip = [];
            editAdmin(x);
        }else if(json.status == 403){
            alert(json.message)
        }
    }
    function clearEdit(){
        admin_name_field = "";
        admin_idrule_field = "";
        admin_status_field = "";
        admin_create_field = "";
        admin_update_field = "";
    }
    initapp();
</script>

{#if sData == "" && admin_username == ""}
    {#if akses_page == true}
        <Home
            on:handleRefreshData={handleRefreshData}
            on:handleEditData={handleEditData}
            {token}
            {admin_listrule}
            {listAdmin}
            {totalrecord}
        />
    {/if}
{:else}
    <Entry
        on:handleBackHalaman={handleBackHalaman}
        on:handleDeleteIplist={handleDeleteIplist}
        on:handleRefreshIplist={handleRefreshIplist}
        {sData}
        {token}
        {admin_listrule}
        {admin_listip}
        {admin_username}
        {admin_name_field}
        {admin_type_field}
        {admin_idrule_field}
        {admin_status_field}
        {admin_create_field}
        {admin_update_field}
    />
{/if}
