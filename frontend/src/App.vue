<template>
    <RootNavbar />
    <router-view />
</template>

<script>
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap/dist/js/bootstrap";
import RootNavbar from "./components/RootNavBar";

export default {
    name: "App",
    components: {
        RootNavbar,
    },
};
</script>

<style></style>
