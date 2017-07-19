type SECTION struct
{
    Id int `db:"id"`;
    Number int `db:"number"`;
    Name string `db:"name"`;
    Text string `db:"text"`;
    Image string `db:"image"`;
}

// ~~

type USER struct
{
    Id int `db:"id"`;
    Email string `db:"email"`;
    Pseudonym string `db:"pseudonym"`;
    Password string `db:"password"`;
    ItIsAdministrator bool `db:"it_is_administrator"`;
}

// ~~

type ARTICLE struct
{
    Id int `db:"id"`;
    Title string `db:"title"`;
    Text string `db:"text"`;
    Image string `db:"image"`;
    Date string `db:"date"`;
    SectionId int `db:"section_id"`;
    UserId int `db:"user_id"`;
    Section * SECTION;
    User * USER;
}

// ~~

type COMMENT struct
{
    Id int `db:"id"`;
    Text string `db:"text"`;
    Date string `db:"date"`;
    ArticleId int `db:"article_id"`;
    UserId int `db:"user_id"`;
    Article * ARTICLE;
    User * USER;
}

// ~~

type SUBSCRIBER struct
{
    Id int `db:"id"`;
    Name string `db:"name"`;
    Email string `db:"email"`;
}

// ~~

func FillStore()
{
    AddSection(
        SECTION
        {
            Id : 1,
            Number : 71,
            Name : "Aririev",
            Text : "Fear did not presume upon a robe and asvajit met them on her hands. Between which left her own cure him but by inner eye. Rolled his pillow bathe green great his heaven changed to troubled years decline still cold lips, full and blood of vice themselves buried beneath their accustomed rounds. Thyself this peaceful spot where peace the livelong day. There great the steeds of softest black eyed have been wrought.",
            Image : "athoucim.png"
        }
        );

    AddSection(
        SECTION
        {
            Id : 2,
            Number : 658,
            Name : "Mareanok",
            Text : "Alone! to pour some seeming aged king shall wander through folly's devious ways are unloosed, shrunk from devas' hovering near from eye checked them, up their treacherous way is richly spread where loathsome worm proclaims a brilliant bloom to despair. Shaded walks winding up with gold richly spread not far ahead. Joined the public view the screen a mountain goat, do as king to part exact in mummy case three thousand varied bloom, I've seen and wine is sent the brotherhood of which taught how sorrow was turned gray. Displaying and comfort men and hands stretched aloft, planted in them all conquering love makes my lonely age of care come the strong.",
            Image : "alladdiz.jpg"
        }
        );

    AddSection(
        SECTION
        {
            Id : 3,
            Number : 316,
            Name : "Arlance",
            Text : "Cry only food with her friends a fitter place where death ne'er by his lair, born a giant cedar's spreading the knee the belt. Asokas one morning found kasyapa full orbed that brahma's rest heart whole, strangers from brahma leads to learn every tongue and child within above, dust amid the common food sagata's daughters brought. Sweet repose in houses rich abundance and halo crowned, stirs all consciousness of iron fly the judges' seat of perspiration on innocence. Fade from chaos reigned and worst of distant realm, gropes blindly they part of ropes swung high. Unlike that saraputra sent up waters roaring far away.",
            Image : "meladadi.png"
        }
        );

    AddSection(
        SECTION
        {
            Id : 4,
            Number : 228,
            Name : "Iharaj",
            Text : "Siddartha from mind from opening chasms threatened to find. From tender flesh is heard the wars and wine is soft cashmere wool, sails have his kingdom and breathe that rise. Humbly entreating him to suffer sorrow fills homes with rarest flowers or the grove, 'let us seem so soft lips look in sacred shrines. Waiting and chandalas clothed with prizes doubly blest can I freely soar, would gladly gave and remove the roads from gross material things.",
            Image : "donenbry.gif"
        }
        );

    AddSection(
        SECTION
        {
            Id : 5,
            Number : 870,
            Name : "Ellemo",
            Text : "Spreading plain sent up strife and bitterness was firm his every living memory behind, circled thrice the parts and calls all earnest seemed rounded perfect rest house all righteousness. Stately swans sailing on rubbish heaps that all and draped with rare gifts returned, preserving recreating day shall battle back my brother prince its base desires. Darkening whole heart yearning for vultures and creeds.",
            Image : "ingcheta.jpg"
        }
        );

    AddSection(
        SECTION
        {
            Id : 6,
            Number : 506,
            Name : "Eeannan",
            Text : "Taught how quick as thick as I but o love! how selfishness bred hypocrisy pretense deceit, free me a forest seemed ever deepening shade. Sharing griefs that brought food with kind heart would end. Waste not neither lack of rose and decks the silver threads, who tunes the singers chanting a sudden flame.",
            Image : "cacehatu.gif"
        }
        );

    AddSection(
        SECTION
        {
            Id : 7,
            Number : 531,
            Name : "Beloodri",
            Text : "Dependents whom few are left not steal not tell, belched forth a sage brahmans and strength was idle crowd. Seeking bliss for waiting prince fell and mount he led. Perforce he answered by eternal perfect love and rejecting mind has come, coming dawn and sees one was woven webs of sorrow die, rode forth an hour by lusts and shunned by acts in their hymns.",
            Image : "toranesu.gif"
        }
        );

    AddSection(
        SECTION
        {
            Id : 8,
            Number : 20,
            Name : "Anamat",
            Text : "Alarmed at all and song's of fields can breathe their robe of pain. Crushing the pearly tears of revelry had heard, spoke and up they circle round the altar rose trees with densest shade. Clapping her husband and lofty hill sank to king of gems to old king shall lead, absorbed at once dedicate to act of whiter skin and creed are unloosed, lascivious lust taste not war to india's farthest bounds freely soar. Teaching the moguls centuries after him turned in power.",
            Image : "endelmid.png"
        }
        );

    AddSection(
        SECTION
        {
            Id : 9,
            Number : 739,
            Name : "Oblesu",
            Text : "Fade from land and brought their muscles seem. Wounds their dead on end in danger desolation death forgot, shrunk from chaos shone upon the panting goat. God made palaces and smiles and asvajit met nature's open war. Prostrate people to speed whose mighty voice he let his cup of ours.",
            Image : "ebhowide.gif"
        }
        );

    AddSection(
        SECTION
        {
            Id : 10,
            Number : 557,
            Name : "Malastivo",
            Text : "Quicker than costly pile was fled amazed nor poisonous fruits may rise, as rahula lay wrestling in even age a garden's wall of life be the brazen sky. Perfect all will and steady but great horns whose wasting form that day to explain, asked him and glorifies the messengers from phalgu's valley opening spring time or terrify. Transfigured seemed deeper skill of cymbals sound body and tender ties, glittering like guilty sleep for said give in dark again.",
            Image : "soseldag.jpg"
        }
        );

    AddUser(
        USER
        {
            Id : 1,
            Email : "shelby.genet@gmail.com",
            Pseudonym : "divicil",
            Password : "aurilukow872",
            ItIsAdministrator : false
        }
        );

    AddUser(
        USER
        {
            Id : 2,
            Email : "tunlin.harsch@gmail.com",
            Pseudonym : "aryellen",
            Password : "desasqeri339",
            ItIsAdministrator : true
        }
        );

    AddUser(
        USER
        {
            Id : 3,
            Email : "erica.lauriston@hotmail.com",
            Pseudonym : "ickartem",
            Password : "zybalanuco50",
            ItIsAdministrator : false
        }
        );

    AddUser(
        USER
        {
            Id : 4,
            Email : "heleen.curley@outlook.com",
            Pseudonym : "gearatol",
            Password : "emmuelip404",
            ItIsAdministrator : true
        }
        );

    AddUser(
        USER
        {
            Id : 5,
            Email : "henka.chun@live.com",
            Pseudonym : "neanarla",
            Password : "argupavu670",
            ItIsAdministrator : true
        }
        );

    AddUser(
        USER
        {
            Id : 6,
            Email : "crystie.benne@gmail.com",
            Pseudonym : "ehritia",
            Password : "urchaguena609",
            ItIsAdministrator : false
        }
        );

    AddUser(
        USER
        {
            Id : 7,
            Email : "helsa.rychlicki@live.com",
            Pseudonym : "eshellta",
            Password : "lisehlal889",
            ItIsAdministrator : true
        }
        );

    AddUser(
        USER
        {
            Id : 8,
            Email : "lachu.obrecht@live.com",
            Pseudonym : "inalisy",
            Password : "anorsisyu386",
            ItIsAdministrator : true
        }
        );

    AddUser(
        USER
        {
            Id : 9,
            Email : "elane.omalley@live.com",
            Pseudonym : "ourgiwie",
            Password : "andelvicho814",
            ItIsAdministrator : false
        }
        );

    AddUser(
        USER
        {
            Id : 10,
            Email : "vax.whitsell@gmail.com",
            Pseudonym : "doluci",
            Password : "lunuwegod782",
            ItIsAdministrator : false
        }
        );

    AddArticle(
        ARTICLE
        {
            Id : 1,
            Title : "Quick its site in heaven.",
            Text : "Good an endless golden age stops short journey taking needful rest, hanging from deepest love whose outer court with green with eagerness. Bent loving eyes are our own sure as clean and e'en while love illumined truth, crying they searched out its mansion to other hands, new quick with sounds almost seemed good the evening shades descend. Pause not so eager to realms of righteousness! hail! great horns, eager crowd gathered about his webs of home. Their sweet jasmine vines with smiles the thriftless multitude the silver light he saw.",
            Image : "uppofari.png",
            Date : "2005-12-18 10:52:57",
            SectionId : 2,
            UserId : 8
        }
        );

    AddArticle(
        ARTICLE
        {
            Id : 2,
            Title : "Instant fall the rocks.",
            Text : "Profoundly learned from real brahman choir responsive to strive to daily life. Secrets are wise else by furious tempests swept, dance in little we pass by discipline and spires and base desires like swine to each. Divides him tremble quiver glow each pleasure when her and towers were temples reared. Driving away with plots and bitterness and west and sleeping world, brushes away but over hills upward devas sweetly as this.",
            Image : "untrulve.gif",
            Date : "2007-05-28 14:52:06",
            SectionId : 3,
            UserId : 6
        }
        );

    AddArticle(
        ARTICLE
        {
            Id : 3,
            Title : "Join returning pilgrim bands.",
            Text : "Bring his way soiled feet unshod and pilgrims countless as in egypt's tombs, uprooting cherished evils one rudely charging him speculations fables all on her eager steeds. Securely wrapped world sad mood he departs with silks and gaunt, both lithe and ate he forgot the kalpa's end, unnumbered funeral train to trim their ablutions made nature's every ill fed many acolytes. Emblem of manliness and clinched so rich man's advocate and slow.",
            Image : "iehlowle.gif",
            Date : "2010-12-19 14:43:44",
            SectionId : 9,
            UserId : 3
        }
        );

    AddArticle(
        ARTICLE
        {
            Id : 4,
            Title : "Nostrils distended wide open wide.",
            Text : "Helpless waifs who has made him from below. Enticing men can cull the bounties free not a path since has fled amazed, sprang to tempt the sounds almost leap down took root. Indulge no trumpeting of war and aid the flames each passer by distant fields, harangue the diligent are vain and catch one leader one night the upward climbed. Harsh words of slender lattice work and sweep smoothly on raven's wings, goats climb and manhood the desert home leave the pleasures please. Manhood and spreading plain comparing great snow a selfish half its wants were suburb huts, clave to sadness joy or yasodhara came near the trees. Often watched by matter's blinding sand storm changing day the romans loved.",
            Image : "alelacop.gif",
            Date : "2014-03-15 01:43:04",
            SectionId : 9,
            UserId : 4
        }
        );

    AddArticle(
        ARTICLE
        {
            Id : 5,
            Title : "Sprang to pluck the stages of lust.",
            Text : "Traders cross currents that inner world that changeless day. Yoked to return to cheer my dwelling near siddartha rise! go they reach the mother's love. Bowed and pruned spreading plains gilded and pure, eight times and said rise in tattered robes. Scarce raised his rays that crush and dim and soothed him pause.",
            Image : "onsaratu.gif",
            Date : "2008-08-03 06:52:28",
            SectionId : 10,
            UserId : 5
        }
        );

    AddArticle(
        ARTICLE
        {
            Id : 6,
            Title : "Weaving with vaulted roofs.",
            Text : "Bade the o'erwrought prince they part of prophets open groves and breathing sweet sleep, come or eyes that folds seemed the bleat of discord sounds of saving men, nor yet ever knew not steal not make. Strange thing is truth his rings and lions springing on himalaya's chain, give him spare my dwelling near some heartache helping hands when fit the depths was passed. Oft have your love alone a lesser strength, perhaps your country and labor ere he gently said with timid pheasants brought their varied song. Gasping for tears streaming from heaven than mortal eye and weary feet through many names, absorbed in after trouble! rest heart and freeze and faint rays to teach, shrink from every recess lighting a sober friend of unrequited love their lover friend!.",
            Image : "achadaly.gif",
            Date : "2013-05-07 21:38:47",
            SectionId : 3,
            UserId : 10
        }
        );

    AddArticle(
        ARTICLE
        {
            Id : 7,
            Title : "Trapusha one is doubly blest.",
            Text : "Who will I gain an if with sacred prophets grew. Asked his and peoples of sick chief ananda the athletes' prizes won, beginning on wave on unfolded to one a star. Six yellow crimson deep in tongue that every taste human limbs, than costly gems attended men with arm hung quivering in dark lurid flames.",
            Image : "ingarmos.jpg",
            Date : "2000-01-09 19:27:50",
            SectionId : 6,
            UserId : 1
        }
        );

    AddArticle(
        ARTICLE
        {
            Id : 8,
            Title : "Spreading branches drooping branches interlocked above.",
            Text : "Lightening each house is doubly blest can ne'er such vain and stunted trees bore yasodhara, told as devadatta side that another glorious light and sleeping city crumbles to want. Joined and heartless selfishness was an outbirth then two their place another glorious flag, swept o'er mountain sea of nature's heights with home leave all. Clearer to fade from star to treat with fountains cooled, transfigured seemed and pine in manhood's prime minister and greed. When you see eye was borne forward on innocence, thought that every walk god's own revile no noxious weeds and dead world. Kill not nothing worth in perfect rest the terraced hill, wary and king with smiles as he moved about his brother buddha said.",
            Image : "ramirvic.gif",
            Date : "2014-05-06 16:59:09",
            SectionId : 7,
            UserId : 8
        }
        );

    AddArticle(
        ARTICLE
        {
            Id : 9,
            Title : "Bent his brother's path and tree to small.",
            Text : "Smothered with cruel kite that hence should the shepherd groping on sphere, weeping they took a rift to ocean's foam. Loaded the tiny sparrow when corrupt declining rome compelled those grim conspiracy, eat the pompous rites ordained by sorrow groping world within their guide familiar with weeds. 'who ever vomits from thy mighty force through every step. Gasping for winter to buy or heart of sparkling joy or wild fear. Dreams disturb the christian's hope of iron fly.",
            Image : "elikasen.gif",
            Date : "2004-07-26 04:32:20",
            SectionId : 10,
            UserId : 7
        }
        );

    AddArticle(
        ARTICLE
        {
            Id : 10,
            Title : "Hand each arrow's flight.",
            Text : "Wonder that sailor's bark or wild boar or in white scarlet yellow robe, tinges with chains of care for their rich attire to savage winter with joy. Birds and stiffened limbs and brotherhood of lesser strength is strange. Six noble men high above a bright presence in hand till some cruel world!, grows into the lamp of clanking chains that hides all would hide. Goats climb the snows on power that these children of rising in summer's sunny hours. Her for all your thoughts like painted butterflies before their steeds harnessed to son, was death as spotless wool or savior lover tasting love's quick sight and tears.",
            Image : "elinnego.png",
            Date : "2014-01-12 12:09:58",
            SectionId : 9,
            UserId : 6
        }
        );

    AddComment(
        COMMENT
        {
            Id : 1,
            Text : "Good the many weary way their law the greatest kings. None with orchards bend beneath whose eyes glittering palace through buddha's use as numberless, calm of skill of sparkling dew from earth's sad for weary hours. Gilding the babylonian and silvery light siddartha gone out, beneath an equal love clearer sounds she reached nirvana's bliss, measure of people still seem so light nearer comes. Where moths ne'er such royal road to blow, morn noon of ancient seed time sends his loved me all devouring sweep down phalgu's vale. From harsh unloving word checked them drove me joyful tidings spread.",
            Date : "2005-10-08 06:40:02",
            ArticleId : 9,
            UserId : 3
        }
        );

    AddComment(
        COMMENT
        {
            Id : 2,
            Text : "Adding each snowy white as sinews of higher worlds below, harsh words from selfishness breeds strife and surging angry stern command of our coarse thoughts, slave to land to live when yasodhara unveiled. Feeling for knowledge than half veils and east the sports and brooded o'er mountain pass, another train until his burdened agitated soul or eat the slandering tongue that humble awe. Rolling hills beyond his eye to troubled years united voice may gain. They go where'er the banyan's columned vaulted shade of early manhood's prime, cry and watch as darts the pillow bathe within themselves in hand.",
            Date : "2007-07-13 23:43:46",
            ArticleId : 7,
            UserId : 6
        }
        );

    AddComment(
        COMMENT
        {
            Id : 3,
            Text : "Tinges the air with leaves covered from ganges' broad domains rich tyrian silk, crying most ancient sages taught each seeming childlike past. Forthwith sudata bowed in air of flying feet and slanderous speech, asokas one for though they saw siddartha's bride. Yes! I gain the links of everything that tower saw.",
            Date : "2001-11-12 06:03:35",
            ArticleId : 10,
            UserId : 3
        }
        );

    AddComment(
        COMMENT
        {
            Id : 4,
            Text : "Wearied and snaps in with closely shaven head as her inmost soul. Seven days with gardens know no altar stood, ruffled the clattering car through this rust and air, kiss its fatal chain of outward things could still would gently swaying gait. Securely wrapped in life where no bird and selfish thoughts, clothing the goal who out their ease the keenest pleasure luxury and still. Transfigured seemed to eye was revealed the charioteers in need not knowing well nigh spent. Weaving with like ours full at siddartha's natal day drew to purity and spontaneous grow, guiding each howdah gaily decked with lightly danced. Two brahmans for thievish greed the sun's slant rays.",
            Date : "2008-08-25 21:24:06",
            ArticleId : 2,
            UserId : 7
        }
        );

    AddComment(
        COMMENT
        {
            Id : 5,
            Text : "Fly from whence a great for sleep in gardens of grandeur rise, they brought forth an angel clothed in him saraputra after swept, fair as they knew so wondrous wise else impaled before that hall. Cringes in time has made him tell them is used and bamboo plumes darkened by night. Circled thrice swept through vacant for heavenly rest, defeated not clogged by distant lands o'er desert wastes.",
            Date : "2002-02-14 05:37:28",
            ArticleId : 6,
            UserId : 3
        }
        );

    AddComment(
        COMMENT
        {
            Id : 6,
            Text : "Seek to lay stretched his noble purposes of present power cast off, each year a chorus of the powers that nature waves of whitest ivory or slaves, bordered by thirst drink for our strength is youth and halo round the pompous funeral piles. Absorbed at such lofty swans sailing on the root. Our muffled heart speaking ne'er escape his unsuspecting prey.",
            Date : "2002-07-13 13:16:24",
            ArticleId : 9,
            UserId : 3
        }
        );

    AddComment(
        COMMENT
        {
            Id : 7,
            Text : "Nature's rich poor waifs who came an infant's smiling face as builders' hammers in smiles. Dispels the twelve mighty and kasyapa joined till coming watch, thought they skim lightly dipping oars and spreading shade his sturdy blows, thick set with plots the giant shadows lay dead to foes. Often stops short sighted men! o faithless priest let every tongue, rolled on an arrow home returned a clearer to watch and foot the grave.",
            Date : "2012-03-16 05:33:58",
            ArticleId : 8,
            UserId : 1
        }
        );

    AddComment(
        COMMENT
        {
            Id : 8,
            Text : "Guarded by india's farthest bounds if led the heart has grown!. Blindly on an indian bride but rose together went, men as brave purna said caught them rise or woe. Diffused a marriage feast three nisaean stock to morrow dust to spare. Sinks in shining waves of cares oppressed with chains that very face.",
            Date : "2011-11-02 19:45:49",
            ArticleId : 1,
            UserId : 8
        }
        );

    AddComment(
        COMMENT
        {
            Id : 9,
            Text : "Her face beaming with this said rise from brahma god their simple free, looked the lowest beasts shall pierce the leader's face away, nostrils distended wide open blinded dazzled eyes yet seeming silver bells. Below he wished to go where'er the worlds as words, one come and decked with sacred peepul's shade his beloved king. Called the pillow bathe in jungles and craft and lions elephants in europe's trackless wastes.",
            Date : "2006-04-11 09:38:12",
            ArticleId : 6,
            UserId : 10
        }
        );

    AddComment(
        COMMENT
        {
            Id : 10,
            Text : "Cherished evils one leader seek a white and grace, lascivious lust of siddartha's birth and righteousness and strong again they placed one palsied limb, tears but selfishness breeds strife and deadly pain. Childhood returned a cloud to jutting cliffs seem yawning chasms seem so like coals of swans, named kantaka pawing and heartless manhood cold succeeds. Only fit the laurel crowned appeared and my arms, twelve strings too rich and oil and livid lips, till he thither yasodhara in even scale the form is ever new. Asokas one leader seek and ward off dust, not injure comfort not from an old believed.",
            Date : "2002-11-19 03:39:34",
            ArticleId : 2,
            UserId : 6
        }
        );

    AddSubscriber(
        SUBSCRIBER
        {
            Id : 1,
            Name : "Arasley",
            Email : "kin.sauder@gmail.com"
        }
        );

    AddSubscriber(
        SUBSCRIBER
        {
            Id : 2,
            Name : "Erdihan",
            Email : "icylyn.raha@yahoo.com"
        }
        );

    AddSubscriber(
        SUBSCRIBER
        {
            Id : 3,
            Name : "Ettaharac",
            Email : "paulinus.mowat@yahoo.com"
        }
        );

    AddSubscriber(
        SUBSCRIBER
        {
            Id : 4,
            Name : "Inafrane",
            Email : "twana.pols@yahoo.com"
        }
        );

    AddSubscriber(
        SUBSCRIBER
        {
            Id : 5,
            Name : "Erlikun",
            Email : "yetty.dsilva@live.com"
        }
        );

    AddSubscriber(
        SUBSCRIBER
        {
            Id : 6,
            Name : "Lonathiko",
            Email : "miranda.damena@gmail.com"
        }
        );

    AddSubscriber(
        SUBSCRIBER
        {
            Id : 7,
            Name : "Uelicfar",
            Email : "asnat.misslitz@live.com"
        }
        );

    AddSubscriber(
        SUBSCRIBER
        {
            Id : 8,
            Name : "Affoskidi",
            Email : "marsha.turney@outlook.com"
        }
        );

    AddSubscriber(
        SUBSCRIBER
        {
            Id : 9,
            Name : "Depidge",
            Email : "gilli.committe@outlook.com"
        }
        );

    AddSubscriber(
        SUBSCRIBER
        {
            Id : 10,
            Name : "Ottlalkag",
            Email : "derrick.steelman@hotmail.com"
        }
        );

}
