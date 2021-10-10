<script>
    import Home from "../adminrule/Home.svelte";
    import Entry from "../adminrule/Entry.svelte";
    let listAdminrule = [];
    let record = "";
    let totalrecord = 0;
    let sData = "";
    let adminrule_id = "";
    let adminrule_name = "";
    let token = localStorage.getItem("token");
    let akses_page = true;
    let adminrule_rule_field = "";
    let adminrule_create_field = "";
    let adminrule_update_field = "";

    const handleBackHalaman = () => {
        adminrule_id = "";
        adminrule_name = "";
        adminrule_create_field = "";
        adminrule_update_field = "";
        sData = "";
        listAdminrule = [];
        handleRefreshData("all");
    };
    const handleRefreshData = (e) => {
        listAdminrule = [];
        totalrecord = 0;
        setTimeout(function () {
            initAdmin();
        }, 500);
    };
    const handleEditData = (e) => {
        adminrule_id = e.detail.e;
        adminrule_name = e.detail.f;
        sData = "Edit";
        editAdmin(adminrule_id);
    };
    async function initapp() {
        const res = await fetch("/api/home", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "ADMINRULE-VIEW",
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
        const res = await fetch("/api/alladminrule", {
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
            let css_status = "";
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listAdminrule = [
                        ...listAdminrule,
                        {
                            adminrule_no: record[i]["adminrule_no"],
                            adminrule_id: record[i]["adminrule_id"],
                            adminrule_nama: record[i]["adminrule_nama"],
                        },
                    ];
                }
            } else {
                alert("Error");
            }
        } else {
            logout();
        }
    }
    async function editAdmin(e) {
        const res = await fetch("/api/editadminrule", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idrule: e,
            }),
        });
        const json = await res.json();
        let record = json.record;

        if (json.status === 400) {
            logout();
        } else {
            for (let i = 0; i < record.length; i++) {
                adminrule_rule_field = record[i]["adminrule_rule"];
                adminrule_create_field = record[i]["adminrule_create"];
                adminrule_update_field = record[i]["adminrule_update"];
            }
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    initapp();
</script>

{#if sData == "" && adminrule_id == ""}
    {#if akses_page == true}
        <Home
            on:handleRefreshData={handleRefreshData}
            on:handleEditData={handleEditData}
            {token}
            {listAdminrule}
            {totalrecord}
        />
    {/if}
{:else}
    <Entry
        on:handleBackHalaman={handleBackHalaman}
        {sData}
        {token}
        {adminrule_id}
        {adminrule_name}
        {adminrule_rule_field}
        {adminrule_create_field}
        {adminrule_update_field}
    />
{/if}
