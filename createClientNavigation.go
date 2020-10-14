package itswizard_siteSetup

import (
	"github.com/itslearninggermany/itszwizard_objects"
	"html/template"
)

func createClientNavi(query string, user itszwizard_objects.SessionUser) template.HTML {

	//CSV
	//Schild
	//Excel
	//Hierarchien

	erg := `<div class="hero-head">
    <a role="button" class="navbar-burger" data-target="navMenu" aria-label="menu" aria-expanded="false">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
    </a>

    <div class="navbar-menu" id="navMenu">
        <!-- Menü -->
        <div class="navbar-end">
            <a class="navbar-item" href="/client/startpage` + query + `">
                Start
            </a>

            <div class="navbar-item has-dropdown is-hoverable">
                <a class="navbar-link">
                    Nutzerdaten synchronisieren
                </a>

                <div class="navbar-dropdown">
                    <a class="navbar-item" href="/client/usersync/csvchoice` + query + `">
                        CSV-Datei hochladen
                    </a>
					<a class="navbar-item" href="/client/usersync/azureactivedirectory` + query + `">
                        Mit einem Azure Active Directory verbinden
					</a>
                    <a class="navbar-item" href="/client/usersync/schildnrw` + query + `">
                        Datei aus Schild-NRW hochladen
                    </a>
					<!--
                    <a class="navbar-item" href="/client/usersync/excelchoice` + query + `">
                        Excel-Datei hochladen
                    </a>
                    <a class="navbar-item" href="/client/showUsers` + query + `">
                        Aktuelle Nutzerdaten
					</a>
                    <a class="navbar-item" href="/client/showGroups` + query + `">
                        Aktuelle Hierarchien bearbeiten
                    </a>
                    <hr class="navbar-divider">
                    <a class="navbar-item" href="/client/usersync/start` + query + `">
                        Synchronisieren
                    </a>  -->
                </div>

            </div>

            <div class="navbar-item has-dropdown is-hoverable">
                <a class="navbar-link">
                    Nutzerdaten herunterladen
                </a>

                <div class="navbar-dropdown">
                    <a class="navbar-item" href="/client/download/excel` + query + `">
                        Als Exceldatei
                    </a>
                </div>
            </div>

            <div class="navbar-item has-dropdown is-hoverable">
                <a class="navbar-link">
                    Nutzereinstellungen und Logs
                </a>

                <div class="navbar-dropdown">
                    <a class="navbar-item" href="/client/options` + query + `">
                        Persönliche Daten
                    </a>
                    <a class="navbar-item" href="/client/changepw` + query + `">
                        Passwort ändern
                    </a>
					<a class="navbar-item" href="/client/log` + query + `">
                        Log
                    </a>
                    <hr class="navbar-divider">
                    <a class="navbar-item" href="/client/logOut` + query + `">
                        Log Out
                    </a>
                </div>
            </div>
        </div>
    </div>
</div>



<script>
    document.addEventListener('DOMContentLoaded', () => {

        // Get all "navbar-burger" elements
        const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);

        // Check if there are any navbar burgers
        if ($navbarBurgers.length > 0) {

            // Add a click event on each of them
            $navbarBurgers.forEach( el => {
                el.addEventListener('click', () => {

                    // Get the target from the "data-target" attribute
                    const target = el.dataset.target;
                    const $target = document.getElementById(target);

                    // Toggle the "is-active" class on both the "navbar-burger" and the "navbar-menu"
                    el.classList.toggle('is-active');
                    $target.classList.toggle('is-active');

                });
            });
        }

    });
</script>`

	return template.HTML(erg)

}
