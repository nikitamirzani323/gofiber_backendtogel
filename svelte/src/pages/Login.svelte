<script>
    import { Container, Row, Col, Button } from "sveltestrap";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";

    let client_ipaddress = "";
    let client_timezone = "";

    const schema = yup.object().shape({
        username: yup.string().required().matches(/^[a-zA-z0-9]+$/, "Username must Character A-Z or a-z or 1-9 ").min(4).max(30),
        password: yup.string().required().min(4).max(50)
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            username: "",
            password: ""
        },
        validationSchema: schema,
        onSubmit:(values) => {
           handleSave(values.username,values.password)
        }
    })

    async function handleSave(username,password) {
        const res = await fetch("/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                username: username,
                password: password,
                ipaddress: client_ipaddress,
                timezone: client_timezone,
            }),
        });
        const json = await res.json();

        if (json.status == 400) {
            alert(json.message);
            $form.username = ""
            $form.password = ""
        } else {
            localStorage.setItem("token", json.token);
            window.location.href = "/";
        }
    }
    async function initTimezone() {
        const res = await fetch("api/healthz");
        if (!res.ok) {
            const message = `An error has occured: ${res.status}`;
            throw new Error(message);
        } else {
            const json = await res.json();
            client_ipaddress = json.real_ip;
            client_timezone = "Asia/Jakarta";
        }
    }
    initTimezone();
    $:{
        if ($errors.username || $errors.password){
            alert($errors.username+"\n"+$errors.password)
            $form.username = ""
            $form.password = ""
        }
    }
</script>

<Container style="margin-top: 70px;">
    <Row>
        <Col md={{size:4, offset:4}} xs="3">
            <div class="card" style="border-radius:0px;margin-top:10px;">
                <div class="card-header" style="background-color: #1f2937;">
                    <span
                        style="color:white;font-weight: bold;font-size: 15px;"
                        >LOGIN</span
                    >
                </div>
                <div class="card-body">
                    <div class="mb-3">
                        <input
                            on:change="{handleChange}"
                            bind:value={$form.username}
                            invalid={$errors.username.length > 0}
                            type="text"
                            class="form-control"
                            placeholder="Username"
                            aria-label="Username"
                        />
                    </div>
                    <div class="mb-3">
                        <input
                            on:change="{handleChange}"
                            bind:value={$form.password}
                            invalid={$errors.password.length > 0}
                            type="password"
                            class="form-control"
                            placeholder="Password"
                            aria-label="Password"
                        />
                    </div>
                    <div class="mb-3">
                        <Button
                            block
                            color="dark"
                            on:click={() => {
                                handleSubmit();
                            }}>SUBMIT</Button
                        >
                    </div>
                </div>
            </div>
        </Col>
    </Row>
</Container>
