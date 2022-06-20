<template>
    <BasicContent>
        <div class="row justify-content-md-center">
            <div class="col-3">
                <form @submit="login">
                    <div class="mb-3">
                        <label for="username" class="form-label"
                            >username</label
                        >
                        <input
                            v-model="username"
                            type="text"
                            class="form-control"
                            id="username"
                        />
                    </div>
                    <div class="mb-3">
                        <label for="password" class="form-label"
                            >password</label
                        >
                        <input
                            v-model="password"
                            type="password"
                            class="form-control"
                            id="password"
                        />
                    </div>
                    <div class="error-message">{{ errorMessage }}</div>
                    <button type="submit" class="btn btn-primary">login</button>
                </form>
            </div>
        </div>
    </BasicContent>
</template>

<script>
import BasicContent from "@/components/BasicContent";
import { ref } from "vue";

export default {
    name: "UserLogin",
    components: {
        BasicContent,
    },

    setup() {
        let username = ref("");
        let password = ref("");
        let errorMessage = ref("");

        const login = () => {
            console.log(username.value, password.value);
        };

        return {
            username,
            password,
            errorMessage,
            login,
        };
    },
};
</script>

<style scoped>
button {
    width: 100%;
}
.error-message {
    color: red;
}
</style>
