<script>
    // components for this layout
    import AdminNavbar from "$lib/components/Navbars/AdminNavbar.svelte";
    import Sidebar from "$lib/components/Sidebar/Sidebar.svelte";
    import HeaderStats from "$lib/components/Headers/HeaderStats.svelte";
    import FooterAdmin from "$lib/components/Footers/FooterAdmin.svelte";
    
    let { children } = $props();
  </script>
  
  <div class="flex-row">
    <Sidebar/>
    <div class="md:h-screen overflow-auto relative md:ml-64 bg-blueGray-100 flex flex-col flex-1">
      <AdminNavbar />
      <div class="flex-1 overflow-auto sm:p-stop-2 sm:pt-16 md:pt-0">
        <div class="px-4 md:px-10 mx-auto w-full mt-10">
          {@render children()}
          <FooterAdmin />
        </div>
      </div>
    </div>
  </div>
  