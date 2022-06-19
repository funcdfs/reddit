<template>
    <nav
        class="navbar navbar-expand-lg navbar-light"
        style="background-color: #e3f2fd"
    >
        <div class="container">
            <a class="navbar-brand" href="#">Home</a>
            <button
                class="navbar-toggler"
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#navbarText"
                aria-controls="navbarText"
                aria-expanded="false"
                aria-label="Toggle navigation"
            >
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarText">
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                    <li class="nav-item">
                        <div class="btn-group">
                            <button
                                type="button"
                                class="btn btn-info dropdown-toggle"
                                data-bs-toggle="dropdown"
                                aria-expanded="false"
                            >
                                Group
                            </button>
                            <ul class="dropdown-menu">
                                <li>
                                    <a class="dropdown-item" href="#">Golang</a>
                                </li>
                                <li>
                                    <a class="dropdown-item" href="#">C++</a>
                                </li>
                                <li>
                                    <a class="dropdown-item" href="#"
                                        >JavaScript</a
                                    >
                                </li>
                                <li><hr class="dropdown-divider" /></li>
                                <li>
                                    <a class="dropdown-item" href="#"
                                        >New Group</a
                                    >
                                </li>
                            </ul>
                        </div>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Friends</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Chat</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Space</a>
                    </li>
                    <li class="nav-item">
                        <form class="d-flex" role="search">
                            <input
                                class="form-control me-2"
                                type="search"
                                placeholder="search something"
                                aria-label="Search"
                            />
                            <button
                                class="btn btn-outline-success"
                                type="submit"
                            >
                                Search
                            </button>
                        </form>
                    </li>
                </ul>
                <span class="navbar-text">
                    <div class="btn-group">
                        <button
                            type="button"
                            class="btn btn-info dropdown-toggle"
                            data-bs-toggle="dropdown"
                            aria-expanded="false"
                        >
                            User Options
                        </button>
                        <ul class="dropdown-menu">
                            <li>
                                <a class="dropdown-item" href="#">Sign In</a>
                            </li>
                            <li>
                                <a class="dropdown-item" href="#">Sign Up</a>
                            </li>
                            <li>
                                <a class="dropdown-item" href="#">Sign Out</a>
                            </li>
                            <li><hr class="dropdown-divider" /></li>
                            <li>
                                <a class="dropdown-item" href="#">Settings</a>
                            </li>
                        </ul>
                    </div>
                </span>
            </div>
        </div>
    </nav>
</template>

<script>
export default {
    name: "RootNavbar",
    components: {},
};
</script>

<style scoped></style>
