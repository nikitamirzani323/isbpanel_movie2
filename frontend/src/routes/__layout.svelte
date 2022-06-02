<script>
    import "../tailwind.css";
    import Nav from "../components/navbar.svelte"
    import Footer from "../components/footer.svelte"
</script>
<Nav />
<div class="p-1 max-w-6xl mx-auto">
    <slot />
</div>
<Footer />