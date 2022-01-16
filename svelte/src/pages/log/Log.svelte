<script>
    import Home from "../log/Home.svelte";
    let listHome = [];
    let record = "";
    let totalrecord = 0;
    let sData = "";
    let admin_username = "";
    let token = localStorage.getItem("token");
    let client_key = "";
    let akses_page = true;
    

   
    const handleRefreshData = (e) => {
        listHome = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 500);
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
            initHome();
        }
    }
    async function initHome() {
        const res = await fetch("/api/log", {
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
            let no = 0
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    let note = record[i]["log_note"];
                    let note_result = note.replace("\n", "<br>");
                    listHome = [
                        ...listHome,
                        {
                            log_no: no,
                            log_id: record[i]["log_id"],
                            log_datetime: record[i]["log_datetime"],
                            log_username: record[i]["log_username"],
                            log_page: record[i]["log_page"],
                            log_tipe: record[i]["log_tipe"],
                            log_note: note_result,
                        },
                    ];
                }
            }
        } else {
            logout();
        }
    }
   
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
   
    
    initapp();
</script>

{#if sData == "" && admin_username == ""}
    {#if akses_page == true}
        <Home
            on:handleRefreshData={handleRefreshData}
            {token}
            {listHome}
            {totalrecord}
        />
    {/if}
{/if}
