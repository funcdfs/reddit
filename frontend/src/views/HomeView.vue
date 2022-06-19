<template>
    <div class="home">
        <BasicContent> home page </BasicContent>
    </div>
</template>

<script>
// @ is an alias to /src
import BasicContent from "../components/BasicContent";

export default {
    name: "HomeView",
    components: {
        BasicContent,
    },
};
</script>

<style scoped>
.container {
    margin-top: 20px;
}
</style>
