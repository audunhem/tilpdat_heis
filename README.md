# tilpdat_heis

Funksjonsspesifikasjonen er	gitt	i	prosa-form	i	neste	avsnitt,	og	prosjektets	
oppgaver	vil være:
1. Analyse av	funksjonsspesifikasjonen	(use	case-modell).
2. Design	av	en	overordnet	systemarkitektur (modularisering,	modulgrensesnitt,	
kommunikasjonsmønster/kopling	mellom	modulene).
3. Moduldesign (oppførsel,	indre	struktur	osv.).
4. Implementasjon av	modulene	i	C	på	en	Linux-maskin.
5. Testing	av	modulene hver	for	seg.
6. Testing	av	det	ferdige	styresystemet.

Funksjonsspesifikasjon
Dette	avsnittet	tar	for	seg	funksjonsspesifikasjonen	for	styringssystemet.	Dette	er utgangpunktet	
for	arbeidet	iht.	V-modellen,	og	vil	danne	grunnlaget	for	FAT-kriteriene (Factory	Acceptance	
Test).
Det	er	et	mål	å	gi	heisen	en	oppførsel	som	likner	en	virkelig	heis.	Styringssystemet som	dere	skal	
lage	skal	tilfredsstille	følgende	funksjonelle	krav.

1 Oppstart
Ved	oppstart	kan	heisstolen være	i	en	ikke-definert	tilstand	mellom	to etasjer.	Ikke-definert	
tilstand	betyr	at	det	ikke	er	mulig	for	styringssystemet å	avgjøre	hvor	heisstolen befinner	seg.	
Heisstolen skal	kjøres	til	en	definert	tilstand før	heisen	skal	reagere	på	knappetrykk.	La	
heisstolen kjøre	opp	eller	ned til	den	kommer	til	en	etasje,	og	stopp	der.	
Dersom	heisen	står	i	en	etasje ved	oppstart, er heisen	allerede	i	en	definert	tilstand.
Hvis	heisen	befinner	over	4.	eller	under	1.	etasje, må	den	manuelt	flyttes	før	programmet	startes.

2 Håndtering	av	bestillinger
Det	er	ikke	ønskelig	å	komme	i	en	situasjon	hvor noen	som	står	i	endeetasjene	(første	eller
fjerde)	må	vente	på	en	heis	som aldri	kommer.	Heisen	skal	heller	ikke	stoppe flere	ganger	enn	
nødvendig,	så	hvis	heisen	på	vei	opp	passerer	en	etasje	der	det	kun	er	bestilt	transport	nedover	
(eller	motsatt),	skal	den	ikke	stoppe.	Vi	antar	at	alle	passasjerer	som	venter	på	heisen	i	en	etasje,
går	på	når	heisen	stopper	der.
Eksempel	på	bruk	av	betjeningsknappene:	Hvis	en står	i	2.	etasje	og	ønsker	å	komme	til	4.	etasje,	
vil	en først	trykke	på knappen	OPP i	2.	etasje	på	etasjepanelet	(dette	tilsvarer	at	en	står	i	2.	etasje	
og	signaliserer	at	en	ønsker	å	kjøre oppover).	Når	heisstolen kommer	og	døren	åpnes,	vil	en så	
trykke	på	bestillingsknappen merket	4 på	heispanelet	(dette	tilsvarer	at	en	går	inn	i	heisen	og	
trykker	på	knappen	for	4.	etasje).

3 Bestillingslys	og	etasjelys	
Når	en	for	eksempel	står	utenfor	heisen	i	en etasje	og trykker	OPP eller	NED,	skal	lyset	i	den	
aktuelle	knappen	lyse	inntil	bestillingen er	ekspedert,	dvs.	heisen	er	ankommet	etasjen	og	døren	
er	åpnet.	Det	samme	gjelder	lyset	i	bestillingsknappene	inne	i	heisen.	Når	heisen beveger	seg,
skal	etasjeindikatorene vise	den	siste	etasjen	heisen	var	i.	Er	heisen	f.eks.	mellom	2.	og 3.	etasje	
og	på	vei	oppover,	skal	altså	lampen	som	indikerer	2.	etasje	være tent.

4 Døren
Når	heisen	kommer	til	en	etasje	der	noen	skal	inn	eller	ut, skal	døren	åpnes og	være	åpen	i	tre
sekunder	(dette	indikeres	ved	at	dørlyset	på	heispanelet	er	tent	i	tre	sekunder).	
Heisen	skal	alltid	stå	stille når	døren	er	åpen.	
Når	tre	sekunder har	passert,	skal døren	lukkes	(lyset	slukkes),	og	heisen	skal	så	håndtere	andre	
ubetjente	bestillinger	i	systemet.	Hvis	det	ikke	er noen	ubetjente	bestillinger,	skal	heisen	stå	i	ro	
med	døren	lukket.

5 Obstruksjon
Obstruksjonsbryteren	skal	ikke	ha	noen	innvirkning	på	systemet.

6 Stoppknappen
Når	en	person	trykker	på	stoppknappen, skal	heisen	stoppe	momentant,	og	alle	bestillinger	skal	
slettes. Nye	bestillinger	skal	ikke	registreres	før	stoppknappen	er	sluppet. Når	stoppknappen	er	
sluppet,	skal	heisen	stå	i	ro	inntil	den	har	fått	en	ny	bestilling.
Hvis	heisen	er	i	en	etasje	når	stoppknappen	trykkes,	skal	døren	åpnes. Når	stoppknappen	
slippes,	skal	døren	forbli	åpen	i	tre	sekunder	og	deretter	lukkes.

7 Generelt
Tenk	over	ulike	situasjoner	som	kan	oppstå	og	hvordan	du	ønsker	at heisen	skal	oppføre	seg	
dersom	du	selv	sto	inne	i	heisen.	Systemet	skal være	robust	mot	uforutsette	hendelser	på	den
måten	at	passasjerenes	sikkerhet	ivaretas.	Det	må	for	eksempel	alltid	være mulig	å	komme	inn	
og	ut	av	heisen	dersom	dette	er	forsvarlig,	altså når	heisen	står	stille	i	en	etasje.	Systemet	må	
også	kunne	håndtere	en	hvilken	som	helst	sekvens	av	knappetrykk	(f.eks.	hvis	noen	uforvarende	
lener	seg	mot	heispanelet	og	trykker	inn	alle	etasjeknappene)	uten	at	det	systemet	«krasjer»,	
sikkerheten	svekkes eller	noen	bestillinger	aldri	blir	betjent.	
