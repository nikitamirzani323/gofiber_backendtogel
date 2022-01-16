<script>
  import { link } from "svelte-spa-router";
  import { Button } from "sveltestrap";

  async function handleLogout() {
    localStorage.clear();
    window.location.href = "/";
  }
</script>

<header>
  <!-- Fixed navbar -->
  <nav
    class="navbar navbar-expand-md navbar-dark fixed-top bg-dark"
    style="background-color: #1f2937!important;">
    <div class="container-fluid">
      <a class="navbar-brand" href="/" style="font-size: 13px;">BACKEND AGEN SDSB4D</a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarCollapse"
        aria-controls="navbarCollapse"
        aria-expanded="false"
        aria-label="Toggle navigation">
        <span class="navbar-toggler-icon" />
      </button>
      <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul class="navbar-nav me-auto mb-2 mb-md-0">
          <li class="nav-item">
            <a
              class="nav-link active"
              aria-current="page"
              use:link
              href="/"
              style="font-size: 13px;">DASHBOARD</a >
          </li>
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              id="navbarDropdown"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
              style="font-size: 13px;">
              TOGEL
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
              <li>
                <a
                  class="dropdown-item"
                  use:link
                  href="/periode"
                  style="font-size: 13px;">PERIODE</a>
              </li>
              <li>
                <a
                  class="dropdown-item"
                  use:link
                  href="/prediksi"
                  style="font-size: 13px;">PREDIKSI</a >
              </li>
              <li>
                <a
                  class="dropdown-item"
                  use:link
                  href="/reportwinlose"
                  style="font-size: 13px;">REPORT WINLOSE</a>
              </li>
            </ul>
          </li>
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              id="navbarDropdown"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
              style="font-size: 13px;">
              MASTER
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
              <li>
                <a
                  class="dropdown-item"
                  use:link
                  href="/pasaran"
                  style="font-size: 13px;">PASARAN</a>
                <a
                  class="dropdown-item"
                  use:link
                  href="/log"
                  style="font-size: 13px;">LOG MANAGEMENT</a>
                <a
                  class="dropdown-item"
                  use:link
                  href="/admin"
                  style="font-size: 13px;">ADMIN MANAGEMENT</a>
                <a
                  class="dropdown-item"
                  use:link
                  href="/adminrule"
                  style="font-size: 13px;">ADMIN RULE</a>
              </li>
            </ul>
          </li>
        </ul>
        <form class="d-flex">
          <button
            on:click={() => {
              handleLogout();
            }}
            class="btn btn-warning btn-sm">LOGOUT</button>
        </form>
      </div>
    </div>
  </nav>
</header>
