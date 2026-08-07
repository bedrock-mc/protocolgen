// Code generated from canonical protocol manifest v2. DO NOT EDIT.

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ActorEventType {
    None = 0,
    Jump = 1,
    Hurt = 2,
    Death = 3,
    StartAttacking = 4,
    StopAttacking = 5,
    TamingFailed = 6,
    TamingSucceeded = 7,
    ShakeWetness = 8,
    EatGrass = 10,
    FishHookBubble = 11,
    FishHookFishPos = 12,
    FishHookHookTime = 13,
    FishHookTease = 14,
    SquidFleeing = 15,
    ZombieConverting = 16,
    PlayAmbient = 17,
    SpawnAlive = 18,
    StartOfferFlower = 19,
    StopOfferFlower = 20,
    LoveHearts = 21,
    VillagerAngry = 22,
    VillagerHappy = 23,
    WitchHatMagic = 24,
    FireworksExplode = 25,
    InLoveHearts = 26,
    SilverfishMergeAnim = 27,
    GuardianAttackSound = 28,
    DrinkPotion = 29,
    ThrowPotion = 30,
    PrimeTntCart = 31,
    PrimeCreeper = 32,
    AirSupply = 33,
    DeprecatedAddPlayerLevels = 34,
    GuardianMiningFatigue = 35,
    AgentSwingArm = 36,
    DragonStartDeathAnim = 37,
    GroundDust = 38,
    Shake = 39,
    Feed = 57,
    BabyAge = 60,
    InstantDeath = 61,
    NotifyTrade = 62,
    LeashDestroyed = 63,
    CaravanUpdated = 64,
    TalismanActivate = 65,
    DeprecatedUpdateStructureFeature = 66,
    PlayerSpawnedMob = 67,
    Puke = 68,
    UpdateStackSize = 69,
    StartSwimming = 70,
    BalloonPop = 71,
    TreasureHunt = 72,
    SummonAgent = 73,
    FinishedChargingItem = 74,
    ActorGrowUp = 76,
    VibrationDetected = 77,
    DrinkMilk = 78,
    ShakeWetnessStop = 79,
    KineticDamageDealt = 80,
    HurtWithoutReceivingDamage = 81,
}

impl TryFrom<u8> for ActorEventType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::Jump),
            2 => Ok(Self::Hurt),
            3 => Ok(Self::Death),
            4 => Ok(Self::StartAttacking),
            5 => Ok(Self::StopAttacking),
            6 => Ok(Self::TamingFailed),
            7 => Ok(Self::TamingSucceeded),
            8 => Ok(Self::ShakeWetness),
            10 => Ok(Self::EatGrass),
            11 => Ok(Self::FishHookBubble),
            12 => Ok(Self::FishHookFishPos),
            13 => Ok(Self::FishHookHookTime),
            14 => Ok(Self::FishHookTease),
            15 => Ok(Self::SquidFleeing),
            16 => Ok(Self::ZombieConverting),
            17 => Ok(Self::PlayAmbient),
            18 => Ok(Self::SpawnAlive),
            19 => Ok(Self::StartOfferFlower),
            20 => Ok(Self::StopOfferFlower),
            21 => Ok(Self::LoveHearts),
            22 => Ok(Self::VillagerAngry),
            23 => Ok(Self::VillagerHappy),
            24 => Ok(Self::WitchHatMagic),
            25 => Ok(Self::FireworksExplode),
            26 => Ok(Self::InLoveHearts),
            27 => Ok(Self::SilverfishMergeAnim),
            28 => Ok(Self::GuardianAttackSound),
            29 => Ok(Self::DrinkPotion),
            30 => Ok(Self::ThrowPotion),
            31 => Ok(Self::PrimeTntCart),
            32 => Ok(Self::PrimeCreeper),
            33 => Ok(Self::AirSupply),
            34 => Ok(Self::DeprecatedAddPlayerLevels),
            35 => Ok(Self::GuardianMiningFatigue),
            36 => Ok(Self::AgentSwingArm),
            37 => Ok(Self::DragonStartDeathAnim),
            38 => Ok(Self::GroundDust),
            39 => Ok(Self::Shake),
            57 => Ok(Self::Feed),
            60 => Ok(Self::BabyAge),
            61 => Ok(Self::InstantDeath),
            62 => Ok(Self::NotifyTrade),
            63 => Ok(Self::LeashDestroyed),
            64 => Ok(Self::CaravanUpdated),
            65 => Ok(Self::TalismanActivate),
            66 => Ok(Self::DeprecatedUpdateStructureFeature),
            67 => Ok(Self::PlayerSpawnedMob),
            68 => Ok(Self::Puke),
            69 => Ok(Self::UpdateStackSize),
            70 => Ok(Self::StartSwimming),
            71 => Ok(Self::BalloonPop),
            72 => Ok(Self::TreasureHunt),
            73 => Ok(Self::SummonAgent),
            74 => Ok(Self::FinishedChargingItem),
            76 => Ok(Self::ActorGrowUp),
            77 => Ok(Self::VibrationDetected),
            78 => Ok(Self::DrinkMilk),
            79 => Ok(Self::ShakeWetnessStop),
            80 => Ok(Self::KineticDamageDealt),
            81 => Ok(Self::HurtWithoutReceivingDamage),
            value => Err(value),
        }
    }
}

impl From<ActorEventType> for u8 {
    fn from(value: ActorEventType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ActorLinkType {
    None = 0,
    Riding = 1,
    Passenger = 2,
}

impl TryFrom<u8> for ActorLinkType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::Riding),
            2 => Ok(Self::Passenger),
            value => Err(value),
        }
    }
}

impl From<ActorLinkType> for u8 {
    fn from(value: ActorLinkType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum ActorType {
    Undefined = 1,
    ItemEntity = 64,
    PrimedTnt = 65,
    FallingBlock = 66,
    MovingBlock = 67,
    Experience = 69,
    EyeOfEnder = 70,
    EnderCrystal = 71,
    FireworksRocket = 72,
    FishingHook = 77,
    Chalkboard = 78,
    Painting = 83,
    LeashKnot = 88,
    BoatRideable = 90,
    LightningBolt = 93,
    AreaEffectCloud = 95,
    Balloon = 107,
    Shield = 117,
    Lectern = 119,
    OminousItemSpawner = 145,
    Cushion = 154,
    ChestBoatRideable = 218,
    Mob = 256,
    Npc = 307,
    Agent = 312,
    ArmorStand = 317,
    TripodCamera = 318,
    Player = 319,
    Bee = 378,
    Piglin = 379,
    PiglinBrute = 383,
    Allay = 390,
    PathfinderMob = 768,
    IronGolem = 788,
    SnowGolem = 789,
    WanderingTrader = 886,
    CopperGolem = 916,
    SulfurCube = 921,
    Monster = 2816,
    Creeper = 2849,
    Slime = 2853,
    EnderMan = 2854,
    Ghast = 2857,
    LavaSlime = 2858,
    Blaze = 2859,
    Witch = 2861,
    Guardian = 2865,
    ElderGuardian = 2866,
    Dragon = 2869,
    Shulker = 2870,
    Vindicator = 2873,
    IllagerBeast = 2875,
    EvocationIllager = 2920,
    Vex = 2921,
    Pillager = 2930,
    ElderGuardianGhost = 2936,
    Warden = 2947,
    Breeze = 2956,
    Creaking = 2962,
    Animal = 4864,
    Chicken = 4874,
    Cow = 4875,
    Pig = 4876,
    Sheep = 4877,
    MushroomCow = 4880,
    Rabbit = 4882,
    PolarBear = 4892,
    Llama = 4893,
    Turtle = 4938,
    Panda = 4977,
    Fox = 4985,
    Hoglin = 4988,
    Strider = 4989,
    Goat = 4992,
    Axolotl = 4994,
    Frog = 4996,
    Camel = 5002,
    Sniffer = 5003,
    Armadillo = 5006,
    HappyGhast = 5011,
    TraderLlama = 5021,
    WaterAnimal = 8960,
    Squid = 8977,
    Dolphin = 8991,
    Pufferfish = 9068,
    Salmon = 9069,
    Tropicalfish = 9071,
    Fish = 9072,
    GlowSquid = 9089,
    Tadpole = 9093,
    Nautilus = 9109,
    TamableAnimal = 21248,
    Wolf = 21262,
    Ocelot = 21270,
    Parrot = 21278,
    Cat = 21323,
    Ambient = 33024,
    Bat = 33043,
    UndeadMonster = 68352,
    PigZombie = 68388,
    WitherBoss = 68404,
    Phantom = 68410,
    Zoglin = 68478,
    CamelHusk = 70552,
    ZombieNautilus = 74646,
    ZombieMonster = 199424,
    Zombie = 199456,
    ZombieVillager = 199468,
    Husk = 199471,
    Drowned = 199534,
    ZombieVillagerV2 = 199540,
    Arthropod = 264960,
    Spider = 264995,
    Silverfish = 264999,
    CaveSpider = 265000,
    Endermite = 265015,
    Minecart = 524288,
    MinecartRideable = 524372,
    MinecartHopper = 524384,
    MinecartTnt = 524385,
    MinecartChest = 524386,
    MinecartFurnace = 524387,
    MinecartCommandBlock = 524388,
    SkeletonMonster = 1116928,
    Skeleton = 1116962,
    Stray = 1116974,
    WitherSkeleton = 1116976,
    Bogged = 1117072,
    Parched = 1117079,
    EquineAnimal = 2118400,
    Horse = 2118423,
    Donkey = 2118424,
    Mule = 2118425,
    SkeletonHorse = 2183962,
    ZombieHorse = 2183963,
    Projectile = 4194304,
    ExperiencePotion = 4194372,
    ShulkerBullet = 4194380,
    DragonFireball = 4194383,
    Snowball = 4194385,
    ThrownEgg = 4194386,
    LargeFireball = 4194389,
    ThrownPotion = 4194390,
    Enderpearl = 4194391,
    WitherSkull = 4194393,
    WitherSkullDangerous = 4194395,
    SmallFireball = 4194398,
    LingeringPotion = 4194405,
    LlamaSpit = 4194406,
    EvocationFang = 4194407,
    IceBomb = 4194410,
    BreezeWindChargeProjectile = 4194445,
    WindChargeProjectile = 4194447,
    AbstractArrow = 8388608,
    Trident = 12582985,
    Arrow = 12582992,
    VillagerBase = 16777984,
    Villager = 16777999,
    VillagerV2 = 16778099,
}

impl TryFrom<i32> for ActorType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            1 => Ok(Self::Undefined),
            64 => Ok(Self::ItemEntity),
            65 => Ok(Self::PrimedTnt),
            66 => Ok(Self::FallingBlock),
            67 => Ok(Self::MovingBlock),
            69 => Ok(Self::Experience),
            70 => Ok(Self::EyeOfEnder),
            71 => Ok(Self::EnderCrystal),
            72 => Ok(Self::FireworksRocket),
            77 => Ok(Self::FishingHook),
            78 => Ok(Self::Chalkboard),
            83 => Ok(Self::Painting),
            88 => Ok(Self::LeashKnot),
            90 => Ok(Self::BoatRideable),
            93 => Ok(Self::LightningBolt),
            95 => Ok(Self::AreaEffectCloud),
            107 => Ok(Self::Balloon),
            117 => Ok(Self::Shield),
            119 => Ok(Self::Lectern),
            145 => Ok(Self::OminousItemSpawner),
            154 => Ok(Self::Cushion),
            218 => Ok(Self::ChestBoatRideable),
            256 => Ok(Self::Mob),
            307 => Ok(Self::Npc),
            312 => Ok(Self::Agent),
            317 => Ok(Self::ArmorStand),
            318 => Ok(Self::TripodCamera),
            319 => Ok(Self::Player),
            378 => Ok(Self::Bee),
            379 => Ok(Self::Piglin),
            383 => Ok(Self::PiglinBrute),
            390 => Ok(Self::Allay),
            768 => Ok(Self::PathfinderMob),
            788 => Ok(Self::IronGolem),
            789 => Ok(Self::SnowGolem),
            886 => Ok(Self::WanderingTrader),
            916 => Ok(Self::CopperGolem),
            921 => Ok(Self::SulfurCube),
            2816 => Ok(Self::Monster),
            2849 => Ok(Self::Creeper),
            2853 => Ok(Self::Slime),
            2854 => Ok(Self::EnderMan),
            2857 => Ok(Self::Ghast),
            2858 => Ok(Self::LavaSlime),
            2859 => Ok(Self::Blaze),
            2861 => Ok(Self::Witch),
            2865 => Ok(Self::Guardian),
            2866 => Ok(Self::ElderGuardian),
            2869 => Ok(Self::Dragon),
            2870 => Ok(Self::Shulker),
            2873 => Ok(Self::Vindicator),
            2875 => Ok(Self::IllagerBeast),
            2920 => Ok(Self::EvocationIllager),
            2921 => Ok(Self::Vex),
            2930 => Ok(Self::Pillager),
            2936 => Ok(Self::ElderGuardianGhost),
            2947 => Ok(Self::Warden),
            2956 => Ok(Self::Breeze),
            2962 => Ok(Self::Creaking),
            4864 => Ok(Self::Animal),
            4874 => Ok(Self::Chicken),
            4875 => Ok(Self::Cow),
            4876 => Ok(Self::Pig),
            4877 => Ok(Self::Sheep),
            4880 => Ok(Self::MushroomCow),
            4882 => Ok(Self::Rabbit),
            4892 => Ok(Self::PolarBear),
            4893 => Ok(Self::Llama),
            4938 => Ok(Self::Turtle),
            4977 => Ok(Self::Panda),
            4985 => Ok(Self::Fox),
            4988 => Ok(Self::Hoglin),
            4989 => Ok(Self::Strider),
            4992 => Ok(Self::Goat),
            4994 => Ok(Self::Axolotl),
            4996 => Ok(Self::Frog),
            5002 => Ok(Self::Camel),
            5003 => Ok(Self::Sniffer),
            5006 => Ok(Self::Armadillo),
            5011 => Ok(Self::HappyGhast),
            5021 => Ok(Self::TraderLlama),
            8960 => Ok(Self::WaterAnimal),
            8977 => Ok(Self::Squid),
            8991 => Ok(Self::Dolphin),
            9068 => Ok(Self::Pufferfish),
            9069 => Ok(Self::Salmon),
            9071 => Ok(Self::Tropicalfish),
            9072 => Ok(Self::Fish),
            9089 => Ok(Self::GlowSquid),
            9093 => Ok(Self::Tadpole),
            9109 => Ok(Self::Nautilus),
            21248 => Ok(Self::TamableAnimal),
            21262 => Ok(Self::Wolf),
            21270 => Ok(Self::Ocelot),
            21278 => Ok(Self::Parrot),
            21323 => Ok(Self::Cat),
            33024 => Ok(Self::Ambient),
            33043 => Ok(Self::Bat),
            68352 => Ok(Self::UndeadMonster),
            68388 => Ok(Self::PigZombie),
            68404 => Ok(Self::WitherBoss),
            68410 => Ok(Self::Phantom),
            68478 => Ok(Self::Zoglin),
            70552 => Ok(Self::CamelHusk),
            74646 => Ok(Self::ZombieNautilus),
            199424 => Ok(Self::ZombieMonster),
            199456 => Ok(Self::Zombie),
            199468 => Ok(Self::ZombieVillager),
            199471 => Ok(Self::Husk),
            199534 => Ok(Self::Drowned),
            199540 => Ok(Self::ZombieVillagerV2),
            264960 => Ok(Self::Arthropod),
            264995 => Ok(Self::Spider),
            264999 => Ok(Self::Silverfish),
            265000 => Ok(Self::CaveSpider),
            265015 => Ok(Self::Endermite),
            524288 => Ok(Self::Minecart),
            524372 => Ok(Self::MinecartRideable),
            524384 => Ok(Self::MinecartHopper),
            524385 => Ok(Self::MinecartTnt),
            524386 => Ok(Self::MinecartChest),
            524387 => Ok(Self::MinecartFurnace),
            524388 => Ok(Self::MinecartCommandBlock),
            1116928 => Ok(Self::SkeletonMonster),
            1116962 => Ok(Self::Skeleton),
            1116974 => Ok(Self::Stray),
            1116976 => Ok(Self::WitherSkeleton),
            1117072 => Ok(Self::Bogged),
            1117079 => Ok(Self::Parched),
            2118400 => Ok(Self::EquineAnimal),
            2118423 => Ok(Self::Horse),
            2118424 => Ok(Self::Donkey),
            2118425 => Ok(Self::Mule),
            2183962 => Ok(Self::SkeletonHorse),
            2183963 => Ok(Self::ZombieHorse),
            4194304 => Ok(Self::Projectile),
            4194372 => Ok(Self::ExperiencePotion),
            4194380 => Ok(Self::ShulkerBullet),
            4194383 => Ok(Self::DragonFireball),
            4194385 => Ok(Self::Snowball),
            4194386 => Ok(Self::ThrownEgg),
            4194389 => Ok(Self::LargeFireball),
            4194390 => Ok(Self::ThrownPotion),
            4194391 => Ok(Self::Enderpearl),
            4194393 => Ok(Self::WitherSkull),
            4194395 => Ok(Self::WitherSkullDangerous),
            4194398 => Ok(Self::SmallFireball),
            4194405 => Ok(Self::LingeringPotion),
            4194406 => Ok(Self::LlamaSpit),
            4194407 => Ok(Self::EvocationFang),
            4194410 => Ok(Self::IceBomb),
            4194445 => Ok(Self::BreezeWindChargeProjectile),
            4194447 => Ok(Self::WindChargeProjectile),
            8388608 => Ok(Self::AbstractArrow),
            12582985 => Ok(Self::Trident),
            12582992 => Ok(Self::Arrow),
            16777984 => Ok(Self::VillagerBase),
            16777999 => Ok(Self::Villager),
            16778099 => Ok(Self::VillagerV2),
            value => Err(value),
        }
    }
}

impl From<ActorType> for i32 {
    fn from(value: ActorType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum AgentActionType {
    Attack = 1,
    Collect = 2,
    Destroy = 3,
    DetectRedstone = 4,
    DetectObstacle = 5,
    Drop = 6,
    DropAll = 7,
    Inspect = 8,
    InspectData = 9,
    InspectItemCount = 10,
    InspectItemDetail = 11,
    InspectItemSpace = 12,
    Interact = 13,
    Move = 14,
    PlaceBlock = 15,
    Till = 16,
    TransferItemTo = 17,
    Turn = 18,
}

impl TryFrom<i32> for AgentActionType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            1 => Ok(Self::Attack),
            2 => Ok(Self::Collect),
            3 => Ok(Self::Destroy),
            4 => Ok(Self::DetectRedstone),
            5 => Ok(Self::DetectObstacle),
            6 => Ok(Self::Drop),
            7 => Ok(Self::DropAll),
            8 => Ok(Self::Inspect),
            9 => Ok(Self::InspectData),
            10 => Ok(Self::InspectItemCount),
            11 => Ok(Self::InspectItemDetail),
            12 => Ok(Self::InspectItemSpace),
            13 => Ok(Self::Interact),
            14 => Ok(Self::Move),
            15 => Ok(Self::PlaceBlock),
            16 => Ok(Self::Till),
            17 => Ok(Self::TransferItemTo),
            18 => Ok(Self::Turn),
            value => Err(value),
        }
    }
}

impl From<AgentActionType> for i32 {
    fn from(value: AgentActionType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum AgentAnimationType {
    ArmSwing = 0,
    Shrug = 1,
}

impl TryFrom<u8> for AgentAnimationType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::ArmSwing),
            1 => Ok(Self::Shrug),
            value => Err(value),
        }
    }
}

impl From<AgentAnimationType> for u8 {
    fn from(value: AgentAnimationType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum AnimateAction {
    NoAction = 0,
    Swing = 1,
    WakeUp = 3,
    CriticalHit = 4,
    MagicCriticalHit = 5,
}

impl TryFrom<u8> for AnimateAction {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::NoAction),
            1 => Ok(Self::Swing),
            3 => Ok(Self::WakeUp),
            4 => Ok(Self::CriticalHit),
            5 => Ok(Self::MagicCriticalHit),
            value => Err(value),
        }
    }
}

impl From<AnimateAction> for u8 {
    fn from(value: AnimateAction) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum AnimationMode {
    None = 0,
    Layers = 1,
    Blocks = 2,
}

impl TryFrom<u8> for AnimationMode {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::Layers),
            2 => Ok(Self::Blocks),
            value => Err(value),
        }
    }
}

impl From<AnimationMode> for u8 {
    fn from(value: AnimationMode) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum BossBarColor {
    Pink = 0,
    Blue = 1,
    Red = 2,
    Green = 3,
    Yellow = 4,
    Purple = 5,
    RebeccaPurple = 6,
    White = 7,
}

impl TryFrom<u8> for BossBarColor {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Pink),
            1 => Ok(Self::Blue),
            2 => Ok(Self::Red),
            3 => Ok(Self::Green),
            4 => Ok(Self::Yellow),
            5 => Ok(Self::Purple),
            6 => Ok(Self::RebeccaPurple),
            7 => Ok(Self::White),
            value => Err(value),
        }
    }
}

impl From<BossBarColor> for u8 {
    fn from(value: BossBarColor) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum BossBarOverlay {
    Progress = 0,
    Notched6 = 1,
    Notched10 = 2,
    Notched12 = 3,
    Notched20 = 4,
}

impl TryFrom<u8> for BossBarOverlay {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Progress),
            1 => Ok(Self::Notched6),
            2 => Ok(Self::Notched10),
            3 => Ok(Self::Notched12),
            4 => Ok(Self::Notched20),
            value => Err(value),
        }
    }
}

impl From<BossBarOverlay> for u8 {
    fn from(value: BossBarOverlay) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum BossEventUpdateType {
    Add = 0,
    PlayerAdded = 1,
    Remove = 2,
    PlayerRemoved = 3,
    UpdatePercent = 4,
    UpdateName = 5,
    UpdateProperties = 6,
    UpdateStyle = 7,
    Query = 8,
}

impl TryFrom<u8> for BossEventUpdateType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Add),
            1 => Ok(Self::PlayerAdded),
            2 => Ok(Self::Remove),
            3 => Ok(Self::PlayerRemoved),
            4 => Ok(Self::UpdatePercent),
            5 => Ok(Self::UpdateName),
            6 => Ok(Self::UpdateProperties),
            7 => Ok(Self::UpdateStyle),
            8 => Ok(Self::Query),
            value => Err(value),
        }
    }
}

impl From<BossEventUpdateType> for u8 {
    fn from(value: BossEventUpdateType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum BuildPlatform {
    Unknown = -1,
    Google = 1,
    IOs = 2,
    Osx = 3,
    Amazon = 4,
    GearVr = 5,
    Uwp = 7,
    Win32 = 8,
    Dedicated = 9,
    TvOs = 10,
    Sony = 11,
    Nx = 12,
    Xbox = 13,
    WindowsPhone = 14,
    Linux = 15,
}

impl TryFrom<i32> for BuildPlatform {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            -1 => Ok(Self::Unknown),
            1 => Ok(Self::Google),
            2 => Ok(Self::IOs),
            3 => Ok(Self::Osx),
            4 => Ok(Self::Amazon),
            5 => Ok(Self::GearVr),
            7 => Ok(Self::Uwp),
            8 => Ok(Self::Win32),
            9 => Ok(Self::Dedicated),
            10 => Ok(Self::TvOs),
            11 => Ok(Self::Sony),
            12 => Ok(Self::Nx),
            13 => Ok(Self::Xbox),
            14 => Ok(Self::WindowsPhone),
            15 => Ok(Self::Linux),
            value => Err(value),
        }
    }
}

impl From<BuildPlatform> for i32 {
    fn from(value: BuildPlatform) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum CameraAimAssistAction {
    Set = 0,
    Clear = 1,
}

impl TryFrom<u8> for CameraAimAssistAction {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Set),
            1 => Ok(Self::Clear),
            value => Err(value),
        }
    }
}

impl From<CameraAimAssistAction> for u8 {
    fn from(value: CameraAimAssistAction) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum CameraAimAssistPresetsPacketOperation {
    Set = 0,
    AddToExisting = 1,
}

impl TryFrom<u8> for CameraAimAssistPresetsPacketOperation {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Set),
            1 => Ok(Self::AddToExisting),
            value => Err(value),
        }
    }
}

impl From<CameraAimAssistPresetsPacketOperation> for u8 {
    fn from(value: CameraAimAssistPresetsPacketOperation) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum CameraAimAssistTargetMode {
    Angle = 0,
    Distance = 1,
}

impl TryFrom<u8> for CameraAimAssistTargetMode {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Angle),
            1 => Ok(Self::Distance),
            value => Err(value),
        }
    }
}

impl From<CameraAimAssistTargetMode> for u8 {
    fn from(value: CameraAimAssistTargetMode) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum CameraAimAssistTargetModeType {
    Angle = 0,
    Distance = 1,
}

impl TryFrom<u8> for CameraAimAssistTargetModeType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Angle),
            1 => Ok(Self::Distance),
            value => Err(value),
        }
    }
}

impl From<CameraAimAssistTargetModeType> for u8 {
    fn from(value: CameraAimAssistTargetModeType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum CameraPresetAudioListener {
    Camera = 0,
    Player = 1,
}

impl TryFrom<u8> for CameraPresetAudioListener {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Camera),
            1 => Ok(Self::Player),
            value => Err(value),
        }
    }
}

impl From<CameraPresetAudioListener> for u8 {
    fn from(value: CameraPresetAudioListener) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum CameraShakeAction {
    Add = 0,
    Stop = 1,
}

impl TryFrom<u8> for CameraShakeAction {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Add),
            1 => Ok(Self::Stop),
            value => Err(value),
        }
    }
}

impl From<CameraShakeAction> for u8 {
    fn from(value: CameraShakeAction) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum CameraShakeType {
    Positional = 0,
    Rotational = 1,
}

impl TryFrom<u8> for CameraShakeType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Positional),
            1 => Ok(Self::Rotational),
            value => Err(value),
        }
    }
}

impl From<CameraShakeType> for u8 {
    fn from(value: CameraShakeType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ChatRestrictionLevel {
    None = 0,
    Dropped = 1,
    Disabled = 2,
}

impl TryFrom<u8> for ChatRestrictionLevel {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::Dropped),
            2 => Ok(Self::Disabled),
            value => Err(value),
        }
    }
}

impl From<ChatRestrictionLevel> for u8 {
    fn from(value: ChatRestrictionLevel) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ClientCameraAimAssistPacketAction {
    SetFromCameraPreset = 0,
    Clear = 1,
}

impl TryFrom<u8> for ClientCameraAimAssistPacketAction {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::SetFromCameraPreset),
            1 => Ok(Self::Clear),
            value => Err(value),
        }
    }
}

impl From<ClientCameraAimAssistPacketAction> for u8 {
    fn from(value: ClientCameraAimAssistPacketAction) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u32)]
pub enum ClientPlayMode {
    Normal = 0,
    Teaser = 1,
    Screen = 2,
    Viewer = 3,
    Reality = 4,
    Placement = 5,
    LivingRoom = 6,
    ExitLevel = 7,
    ExitLevelLivingRoom = 8,
    NumModes = 9,
}

impl TryFrom<u32> for ClientPlayMode {
    type Error = u32;

    fn try_from(value: u32) -> Result<Self, u32> {
        match value {
            0 => Ok(Self::Normal),
            1 => Ok(Self::Teaser),
            2 => Ok(Self::Screen),
            3 => Ok(Self::Viewer),
            4 => Ok(Self::Reality),
            5 => Ok(Self::Placement),
            6 => Ok(Self::LivingRoom),
            7 => Ok(Self::ExitLevel),
            8 => Ok(Self::ExitLevelLivingRoom),
            9 => Ok(Self::NumModes),
            value => Err(value),
        }
    }
}

impl From<ClientPlayMode> for u32 {
    fn from(value: ClientPlayMode) -> Self {
        value as u32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ClientboundTextureShiftAction {
    Invalid = 0,
    Initialize = 1,
    Start = 2,
    SetEnabled = 3,
    Sync = 4,
}

impl TryFrom<u8> for ClientboundTextureShiftAction {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Invalid),
            1 => Ok(Self::Initialize),
            2 => Ok(Self::Start),
            3 => Ok(Self::SetEnabled),
            4 => Ok(Self::Sync),
            value => Err(value),
        }
    }
}

impl From<ClientboundTextureShiftAction> for u8 {
    fn from(value: ClientboundTextureShiftAction) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum CodeBuilderExecutionStateCodeStatus {
    None = 0,
    NotStarted = 1,
    InProgress = 2,
    Paused = 3,
    Error = 4,
    Succeeded = 5,
}

impl TryFrom<u8> for CodeBuilderExecutionStateCodeStatus {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::NotStarted),
            2 => Ok(Self::InProgress),
            3 => Ok(Self::Paused),
            4 => Ok(Self::Error),
            5 => Ok(Self::Succeeded),
            value => Err(value),
        }
    }
}

impl From<CodeBuilderExecutionStateCodeStatus> for u8 {
    fn from(value: CodeBuilderExecutionStateCodeStatus) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum CodeBuilderStorageQueryOptionsCategory {
    None = 0,
    CodeStatus = 1,
    Instantiation = 2,
}

impl TryFrom<u8> for CodeBuilderStorageQueryOptionsCategory {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::CodeStatus),
            2 => Ok(Self::Instantiation),
            value => Err(value),
        }
    }
}

impl From<CodeBuilderStorageQueryOptionsCategory> for u8 {
    fn from(value: CodeBuilderStorageQueryOptionsCategory) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum CodeBuilderStorageQueryOptionsOperation {
    None = 0,
    Get = 1,
    Set = 2,
    Reset = 3,
}

impl TryFrom<u8> for CodeBuilderStorageQueryOptionsOperation {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::Get),
            2 => Ok(Self::Set),
            3 => Ok(Self::Reset),
            value => Err(value),
        }
    }
}

impl From<CodeBuilderStorageQueryOptionsOperation> for u8 {
    fn from(value: CodeBuilderStorageQueryOptionsOperation) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum CommandPermissionLevel {
    Any = 0,
    GameDirectors = 1,
    Admin = 2,
    Host = 3,
    Owner = 4,
    Internal = 5,
}

impl TryFrom<u8> for CommandPermissionLevel {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Any),
            1 => Ok(Self::GameDirectors),
            2 => Ok(Self::Admin),
            3 => Ok(Self::Host),
            4 => Ok(Self::Owner),
            5 => Ok(Self::Internal),
            value => Err(value),
        }
    }
}

impl From<CommandPermissionLevel> for u8 {
    fn from(value: CommandPermissionLevel) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum ConnectionDisconnectFailReason {
    Unknown = 0,
    CantConnectNoInternet = 1,
    NoPermissions = 2,
    UnrecoverableError = 3,
    ThirdPartyBlocked = 4,
    ThirdPartyNoInternet = 5,
    ThirdPartyBadIp = 6,
    ThirdPartyNoServerOrServerLocked = 7,
    VersionMismatch = 8,
    SkinIssue = 9,
    InviteSessionNotFound = 10,
    EduLevelSettingsMissing = 11,
    LocalServerNotFound = 12,
    LegacyDisconnect = 13,
    InternalUserLeaveGameAttempted = 14,
    PlatformLockedSkinsError = 15,
    RealmsWorldUnassigned = 16,
    RealmsServerCantConnect = 17,
    RealmsServerHidden = 18,
    RealmsServerDisabledBeta = 19,
    RealmsServerDisabled = 20,
    CrossPlatformDisabled = 21,
    TestonlyCantConnect = 22,
    SessionNotFound = 23,
    ClientSettingsIncompatibleWithServer = 24,
    ServerFull = 25,
    InvalidPlatformSkin = 26,
    EditionVersionMismatch = 27,
    EditionMismatch = 28,
    LevelNewerThanExeVersion = 29,
    InternalNoFailOccurred = 30,
    BannedSkin = 31,
    Timeout = 32,
    ServerNotFound = 33,
    OutdatedServer = 34,
    OutdatedClient = 35,
    NoPremiumPlatform = 36,
    MultiplayerDisabled = 37,
    NoWiFi = 38,
    WorldCorruption = 39,
    NoReason = 40,
    Disconnected = 41,
    InvalidPlayer = 42,
    LoggedInOtherLocation = 43,
    ServerIdConflict = 44,
    NotAllowed = 45,
    NotAuthenticated = 46,
    InvalidTenant = 47,
    UnknownPacket = 48,
    UnexpectedPacket = 49,
    InvalidCommandRequestPacket = 50,
    HostSuspended = 51,
    LoginPacketNoRequest = 52,
    LoginPacketNoCert = 53,
    MissingClient = 54,
    Kicked = 55,
    KickedForExploit = 56,
    KickedForIdle = 57,
    ResourcePackProblem = 58,
    IncompatiblePack = 59,
    OutOfStorage = 60,
    InvalidLevel = 61,
    DisconnectPacket = 62,
    BlockMismatch = 63,
    InvalidHeights = 64,
    InvalidWidths = 65,
    ConnectionLost = 66,
    ZombieConnection = 67,
    Shutdown = 68,
    ReasonNotSet = 69,
    LoadingStateTimeout = 70,
    ResourcePackLoadingFailed = 71,
    SearchingForSessionLoadingScreenFailed = 72,
    NetherNetProtocolVersion = 73,
    SubsystemStatusError = 74,
    EmptyAuthFromDiscovery = 75,
    EmptyUrlFromDiscovery = 76,
    ExpiredAuthFromDiscovery = 77,
    UnknownSignalServiceSignInFailure = 78,
    XblJoinLobbyFailure = 79,
    UnspecifiedClientInstanceDisconnection = 80,
    NetherNetSessionNotFound = 81,
    NetherNetCreatePeerConnection = 82,
    NetherNetIce = 83,
    NetherNetConnectRequest = 84,
    NetherNetConnectResponse = 85,
    NetherNetNegotiationTimeout = 86,
    NetherNetInactivityTimeout = 87,
    StaleConnectionBeingReplaced = 88,
    RealmsSessionNotFound = 89,
    BadPacket = 90,
    NetherNetFailedToCreateOffer = 91,
    NetherNetFailedToCreateAnswer = 92,
    NetherNetFailedToSetLocalDescription = 93,
    NetherNetFailedToSetRemoteDescription = 94,
    NetherNetNegotiationTimeoutWaitingForResponse = 95,
    NetherNetNegotiationTimeoutWaitingForAccept = 96,
    NetherNetIncomingConnectionIgnored = 97,
    NetherNetSignalingParsingFailure = 98,
    NetherNetSignalingUnknownError = 99,
    NetherNetSignalingUnicastDeliveryFailed = 100,
    NetherNetSignalingBroadcastDeliveryFailed = 101,
    NetherNetSignalingGenericDeliveryFailed = 102,
    EditorMismatchEditorWorld = 103,
    EditorMismatchVanillaWorld = 104,
    WorldTransferNotPrimaryClient = 105,
    InternalRequestServerShutdown = 106,
    ClientGameSetupCancelled = 107,
    ClientGameSetupFailed = 108,
    NoVenue = 109,
    NetherNetSignalingSigninFailed = 110,
    SessionAccessDenied = 111,
    ServiceSigninIssue = 112,
    NetherNetNoSignalingChannel = 113,
    NetherNetNotLoggedIn = 114,
    NetherNetClientSignalingError = 115,
    SubClientLoginDisabled = 116,
    DeepLinkTryingToOpenDemoWorldWhileSignedIn = 117,
    AsyncJoinTaskDenied = 118,
    RealmsTimelineRequired = 119,
    GuestWithoutHost = 120,
    FailedToJoinExperience = 121,
    NetherNetDataChannelClosed = 122,
    DiscoveryEnvironmentMismatch = 123,
    HostWithoutKeys = 124,
    HostSignedOut = 125,
    ScriptWatchdogException = 126,
    ScriptMemoryLimitExceeded = 127,
    StorageLowDuringGameplay = 128,
    StorageFullDuringGameplay = 129,
    LevelStorageCorruption = 130,
    EditionMismatchVanillaToEdu = 131,
    EditionMismatchEduToVanilla = 132,
    EditorMismatchEditorToVanilla = 133,
    EditorMismatchVanillaToEditor = 134,
    DenyListed = 135,
    NonceMissing = 136,
    NonceNotFound = 137,
    NonceExpired = 138,
    NonceNotValid = 139,
    HostDisconnected = 140,
    EditorJoinIntentPolicyFailure = 141,
    NetherNetIdentityNotAllowed = 142,
    InvalidName = 143,
    ExpiredToken = 144,
    HostAcceptsNoTypeOfAuth = 145,
    NotAuthenticatedFastFail = 146,
    EditorNotAllowed = 147,
}

impl TryFrom<i32> for ConnectionDisconnectFailReason {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Unknown),
            1 => Ok(Self::CantConnectNoInternet),
            2 => Ok(Self::NoPermissions),
            3 => Ok(Self::UnrecoverableError),
            4 => Ok(Self::ThirdPartyBlocked),
            5 => Ok(Self::ThirdPartyNoInternet),
            6 => Ok(Self::ThirdPartyBadIp),
            7 => Ok(Self::ThirdPartyNoServerOrServerLocked),
            8 => Ok(Self::VersionMismatch),
            9 => Ok(Self::SkinIssue),
            10 => Ok(Self::InviteSessionNotFound),
            11 => Ok(Self::EduLevelSettingsMissing),
            12 => Ok(Self::LocalServerNotFound),
            13 => Ok(Self::LegacyDisconnect),
            14 => Ok(Self::InternalUserLeaveGameAttempted),
            15 => Ok(Self::PlatformLockedSkinsError),
            16 => Ok(Self::RealmsWorldUnassigned),
            17 => Ok(Self::RealmsServerCantConnect),
            18 => Ok(Self::RealmsServerHidden),
            19 => Ok(Self::RealmsServerDisabledBeta),
            20 => Ok(Self::RealmsServerDisabled),
            21 => Ok(Self::CrossPlatformDisabled),
            22 => Ok(Self::TestonlyCantConnect),
            23 => Ok(Self::SessionNotFound),
            24 => Ok(Self::ClientSettingsIncompatibleWithServer),
            25 => Ok(Self::ServerFull),
            26 => Ok(Self::InvalidPlatformSkin),
            27 => Ok(Self::EditionVersionMismatch),
            28 => Ok(Self::EditionMismatch),
            29 => Ok(Self::LevelNewerThanExeVersion),
            30 => Ok(Self::InternalNoFailOccurred),
            31 => Ok(Self::BannedSkin),
            32 => Ok(Self::Timeout),
            33 => Ok(Self::ServerNotFound),
            34 => Ok(Self::OutdatedServer),
            35 => Ok(Self::OutdatedClient),
            36 => Ok(Self::NoPremiumPlatform),
            37 => Ok(Self::MultiplayerDisabled),
            38 => Ok(Self::NoWiFi),
            39 => Ok(Self::WorldCorruption),
            40 => Ok(Self::NoReason),
            41 => Ok(Self::Disconnected),
            42 => Ok(Self::InvalidPlayer),
            43 => Ok(Self::LoggedInOtherLocation),
            44 => Ok(Self::ServerIdConflict),
            45 => Ok(Self::NotAllowed),
            46 => Ok(Self::NotAuthenticated),
            47 => Ok(Self::InvalidTenant),
            48 => Ok(Self::UnknownPacket),
            49 => Ok(Self::UnexpectedPacket),
            50 => Ok(Self::InvalidCommandRequestPacket),
            51 => Ok(Self::HostSuspended),
            52 => Ok(Self::LoginPacketNoRequest),
            53 => Ok(Self::LoginPacketNoCert),
            54 => Ok(Self::MissingClient),
            55 => Ok(Self::Kicked),
            56 => Ok(Self::KickedForExploit),
            57 => Ok(Self::KickedForIdle),
            58 => Ok(Self::ResourcePackProblem),
            59 => Ok(Self::IncompatiblePack),
            60 => Ok(Self::OutOfStorage),
            61 => Ok(Self::InvalidLevel),
            62 => Ok(Self::DisconnectPacket),
            63 => Ok(Self::BlockMismatch),
            64 => Ok(Self::InvalidHeights),
            65 => Ok(Self::InvalidWidths),
            66 => Ok(Self::ConnectionLost),
            67 => Ok(Self::ZombieConnection),
            68 => Ok(Self::Shutdown),
            69 => Ok(Self::ReasonNotSet),
            70 => Ok(Self::LoadingStateTimeout),
            71 => Ok(Self::ResourcePackLoadingFailed),
            72 => Ok(Self::SearchingForSessionLoadingScreenFailed),
            73 => Ok(Self::NetherNetProtocolVersion),
            74 => Ok(Self::SubsystemStatusError),
            75 => Ok(Self::EmptyAuthFromDiscovery),
            76 => Ok(Self::EmptyUrlFromDiscovery),
            77 => Ok(Self::ExpiredAuthFromDiscovery),
            78 => Ok(Self::UnknownSignalServiceSignInFailure),
            79 => Ok(Self::XblJoinLobbyFailure),
            80 => Ok(Self::UnspecifiedClientInstanceDisconnection),
            81 => Ok(Self::NetherNetSessionNotFound),
            82 => Ok(Self::NetherNetCreatePeerConnection),
            83 => Ok(Self::NetherNetIce),
            84 => Ok(Self::NetherNetConnectRequest),
            85 => Ok(Self::NetherNetConnectResponse),
            86 => Ok(Self::NetherNetNegotiationTimeout),
            87 => Ok(Self::NetherNetInactivityTimeout),
            88 => Ok(Self::StaleConnectionBeingReplaced),
            89 => Ok(Self::RealmsSessionNotFound),
            90 => Ok(Self::BadPacket),
            91 => Ok(Self::NetherNetFailedToCreateOffer),
            92 => Ok(Self::NetherNetFailedToCreateAnswer),
            93 => Ok(Self::NetherNetFailedToSetLocalDescription),
            94 => Ok(Self::NetherNetFailedToSetRemoteDescription),
            95 => Ok(Self::NetherNetNegotiationTimeoutWaitingForResponse),
            96 => Ok(Self::NetherNetNegotiationTimeoutWaitingForAccept),
            97 => Ok(Self::NetherNetIncomingConnectionIgnored),
            98 => Ok(Self::NetherNetSignalingParsingFailure),
            99 => Ok(Self::NetherNetSignalingUnknownError),
            100 => Ok(Self::NetherNetSignalingUnicastDeliveryFailed),
            101 => Ok(Self::NetherNetSignalingBroadcastDeliveryFailed),
            102 => Ok(Self::NetherNetSignalingGenericDeliveryFailed),
            103 => Ok(Self::EditorMismatchEditorWorld),
            104 => Ok(Self::EditorMismatchVanillaWorld),
            105 => Ok(Self::WorldTransferNotPrimaryClient),
            106 => Ok(Self::InternalRequestServerShutdown),
            107 => Ok(Self::ClientGameSetupCancelled),
            108 => Ok(Self::ClientGameSetupFailed),
            109 => Ok(Self::NoVenue),
            110 => Ok(Self::NetherNetSignalingSigninFailed),
            111 => Ok(Self::SessionAccessDenied),
            112 => Ok(Self::ServiceSigninIssue),
            113 => Ok(Self::NetherNetNoSignalingChannel),
            114 => Ok(Self::NetherNetNotLoggedIn),
            115 => Ok(Self::NetherNetClientSignalingError),
            116 => Ok(Self::SubClientLoginDisabled),
            117 => Ok(Self::DeepLinkTryingToOpenDemoWorldWhileSignedIn),
            118 => Ok(Self::AsyncJoinTaskDenied),
            119 => Ok(Self::RealmsTimelineRequired),
            120 => Ok(Self::GuestWithoutHost),
            121 => Ok(Self::FailedToJoinExperience),
            122 => Ok(Self::NetherNetDataChannelClosed),
            123 => Ok(Self::DiscoveryEnvironmentMismatch),
            124 => Ok(Self::HostWithoutKeys),
            125 => Ok(Self::HostSignedOut),
            126 => Ok(Self::ScriptWatchdogException),
            127 => Ok(Self::ScriptMemoryLimitExceeded),
            128 => Ok(Self::StorageLowDuringGameplay),
            129 => Ok(Self::StorageFullDuringGameplay),
            130 => Ok(Self::LevelStorageCorruption),
            131 => Ok(Self::EditionMismatchVanillaToEdu),
            132 => Ok(Self::EditionMismatchEduToVanilla),
            133 => Ok(Self::EditorMismatchEditorToVanilla),
            134 => Ok(Self::EditorMismatchVanillaToEditor),
            135 => Ok(Self::DenyListed),
            136 => Ok(Self::NonceMissing),
            137 => Ok(Self::NonceNotFound),
            138 => Ok(Self::NonceExpired),
            139 => Ok(Self::NonceNotValid),
            140 => Ok(Self::HostDisconnected),
            141 => Ok(Self::EditorJoinIntentPolicyFailure),
            142 => Ok(Self::NetherNetIdentityNotAllowed),
            143 => Ok(Self::InvalidName),
            144 => Ok(Self::ExpiredToken),
            145 => Ok(Self::HostAcceptsNoTypeOfAuth),
            146 => Ok(Self::NotAuthenticatedFastFail),
            147 => Ok(Self::EditorNotAllowed),
            value => Err(value),
        }
    }
}

impl From<ConnectionDisconnectFailReason> for i32 {
    fn from(value: ConnectionDisconnectFailReason) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ContainerEnumName {
    AnvilInputContainer = 0,
    AnvilMaterialContainer = 1,
    AnvilResultPreviewContainer = 2,
    SmithingTableInputContainer = 3,
    SmithingTableMaterialContainer = 4,
    SmithingTableResultPreviewContainer = 5,
    ArmorContainer = 6,
    LevelEntityContainer = 7,
    BeaconPaymentContainer = 8,
    BrewingStandInputContainer = 9,
    BrewingStandResultContainer = 10,
    BrewingStandFuelContainer = 11,
    CombinedHotbarAndInventoryContainer = 12,
    CraftingInputContainer = 13,
    CraftingOutputPreviewContainer = 14,
    RecipeConstructionContainer = 15,
    RecipeNatureContainer = 16,
    RecipeItemsContainer = 17,
    RecipeSearchContainer = 18,
    RecipeSearchBarContainer = 19,
    RecipeEquipmentContainer = 20,
    RecipeBookContainer = 21,
    EnchantingInputContainer = 22,
    EnchantingMaterialContainer = 23,
    FurnaceFuelContainer = 24,
    FurnaceIngredientContainer = 25,
    FurnaceResultContainer = 26,
    HorseEquipContainer = 27,
    HotbarContainer = 28,
    InventoryContainer = 29,
    ShulkerBoxContainer = 30,
    TradeIngredient1Container = 31,
    TradeIngredient2Container = 32,
    TradeResultPreviewContainer = 33,
    OffhandContainer = 34,
    CompoundCreatorInput = 35,
    CompoundCreatorOutputPreview = 36,
    ElementConstructorOutputPreview = 37,
    MaterialReducerInput = 38,
    MaterialReducerOutput = 39,
    LabTableInput = 40,
    LoomInputContainer = 41,
    LoomDyeContainer = 42,
    LoomMaterialContainer = 43,
    LoomResultPreviewContainer = 44,
    BlastFurnaceIngredientContainer = 45,
    SmokerIngredientContainer = 46,
    Trade2Ingredient1Container = 47,
    Trade2Ingredient2Container = 48,
    Trade2ResultPreviewContainer = 49,
    GrindstoneInputContainer = 50,
    GrindstoneAdditionalContainer = 51,
    GrindstoneResultPreviewContainer = 52,
    StonecutterInputContainer = 53,
    StonecutterResultPreviewContainer = 54,
    CartographyInputContainer = 55,
    CartographyAdditionalContainer = 56,
    CartographyResultPreviewContainer = 57,
    BarrelContainer = 58,
    CursorContainer = 59,
    CreatedOutputContainer = 60,
    SmithingTableTemplateContainer = 61,
    CrafterLevelEntityContainer = 62,
    DynamicContainer = 63,
    RecipeFoodContainer = 64,
    RecipeBlocksContainer = 65,
    RecipeFurnaceItemsContainer = 66,
}

impl TryFrom<u8> for ContainerEnumName {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::AnvilInputContainer),
            1 => Ok(Self::AnvilMaterialContainer),
            2 => Ok(Self::AnvilResultPreviewContainer),
            3 => Ok(Self::SmithingTableInputContainer),
            4 => Ok(Self::SmithingTableMaterialContainer),
            5 => Ok(Self::SmithingTableResultPreviewContainer),
            6 => Ok(Self::ArmorContainer),
            7 => Ok(Self::LevelEntityContainer),
            8 => Ok(Self::BeaconPaymentContainer),
            9 => Ok(Self::BrewingStandInputContainer),
            10 => Ok(Self::BrewingStandResultContainer),
            11 => Ok(Self::BrewingStandFuelContainer),
            12 => Ok(Self::CombinedHotbarAndInventoryContainer),
            13 => Ok(Self::CraftingInputContainer),
            14 => Ok(Self::CraftingOutputPreviewContainer),
            15 => Ok(Self::RecipeConstructionContainer),
            16 => Ok(Self::RecipeNatureContainer),
            17 => Ok(Self::RecipeItemsContainer),
            18 => Ok(Self::RecipeSearchContainer),
            19 => Ok(Self::RecipeSearchBarContainer),
            20 => Ok(Self::RecipeEquipmentContainer),
            21 => Ok(Self::RecipeBookContainer),
            22 => Ok(Self::EnchantingInputContainer),
            23 => Ok(Self::EnchantingMaterialContainer),
            24 => Ok(Self::FurnaceFuelContainer),
            25 => Ok(Self::FurnaceIngredientContainer),
            26 => Ok(Self::FurnaceResultContainer),
            27 => Ok(Self::HorseEquipContainer),
            28 => Ok(Self::HotbarContainer),
            29 => Ok(Self::InventoryContainer),
            30 => Ok(Self::ShulkerBoxContainer),
            31 => Ok(Self::TradeIngredient1Container),
            32 => Ok(Self::TradeIngredient2Container),
            33 => Ok(Self::TradeResultPreviewContainer),
            34 => Ok(Self::OffhandContainer),
            35 => Ok(Self::CompoundCreatorInput),
            36 => Ok(Self::CompoundCreatorOutputPreview),
            37 => Ok(Self::ElementConstructorOutputPreview),
            38 => Ok(Self::MaterialReducerInput),
            39 => Ok(Self::MaterialReducerOutput),
            40 => Ok(Self::LabTableInput),
            41 => Ok(Self::LoomInputContainer),
            42 => Ok(Self::LoomDyeContainer),
            43 => Ok(Self::LoomMaterialContainer),
            44 => Ok(Self::LoomResultPreviewContainer),
            45 => Ok(Self::BlastFurnaceIngredientContainer),
            46 => Ok(Self::SmokerIngredientContainer),
            47 => Ok(Self::Trade2Ingredient1Container),
            48 => Ok(Self::Trade2Ingredient2Container),
            49 => Ok(Self::Trade2ResultPreviewContainer),
            50 => Ok(Self::GrindstoneInputContainer),
            51 => Ok(Self::GrindstoneAdditionalContainer),
            52 => Ok(Self::GrindstoneResultPreviewContainer),
            53 => Ok(Self::StonecutterInputContainer),
            54 => Ok(Self::StonecutterResultPreviewContainer),
            55 => Ok(Self::CartographyInputContainer),
            56 => Ok(Self::CartographyAdditionalContainer),
            57 => Ok(Self::CartographyResultPreviewContainer),
            58 => Ok(Self::BarrelContainer),
            59 => Ok(Self::CursorContainer),
            60 => Ok(Self::CreatedOutputContainer),
            61 => Ok(Self::SmithingTableTemplateContainer),
            62 => Ok(Self::CrafterLevelEntityContainer),
            63 => Ok(Self::DynamicContainer),
            64 => Ok(Self::RecipeFoodContainer),
            65 => Ok(Self::RecipeBlocksContainer),
            66 => Ok(Self::RecipeFurnaceItemsContainer),
            value => Err(value),
        }
    }
}

impl From<ContainerEnumName> for u8 {
    fn from(value: ContainerEnumName) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ControlSchemeScheme {
    LockedPlayerRelativeStrafe = 0,
    CameraRelative = 1,
    CameraRelativeStrafe = 2,
    PlayerRelative = 3,
    PlayerRelativeStrafe = 4,
}

impl TryFrom<u8> for ControlSchemeScheme {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::LockedPlayerRelativeStrafe),
            1 => Ok(Self::CameraRelative),
            2 => Ok(Self::CameraRelativeStrafe),
            3 => Ok(Self::PlayerRelative),
            4 => Ok(Self::PlayerRelativeStrafe),
            value => Err(value),
        }
    }
}

impl From<ControlSchemeScheme> for u8 {
    fn from(value: ControlSchemeScheme) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum CoordinateEvaluationOrder {
    Xyz = 0,
    Xzy = 1,
    Yxz = 2,
    Yzx = 3,
    Zxy = 4,
    Zyx = 5,
}

impl TryFrom<i32> for CoordinateEvaluationOrder {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Xyz),
            1 => Ok(Self::Xzy),
            2 => Ok(Self::Yxz),
            3 => Ok(Self::Yzx),
            4 => Ok(Self::Zxy),
            5 => Ok(Self::Zyx),
            value => Err(value),
        }
    }
}

impl From<CoordinateEvaluationOrder> for i32 {
    fn from(value: CoordinateEvaluationOrder) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum CreativeItemCategory {
    Construction = 1,
    Nature = 2,
    Equipment = 3,
    Items = 4,
    ItemCommandOnly = 5,
}

impl TryFrom<u8> for CreativeItemCategory {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            1 => Ok(Self::Construction),
            2 => Ok(Self::Nature),
            3 => Ok(Self::Equipment),
            4 => Ok(Self::Items),
            5 => Ok(Self::ItemCommandOnly),
            value => Err(value),
        }
    }
}

impl From<CreativeItemCategory> for u8 {
    fn from(value: CreativeItemCategory) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum DataItemType {
    Byte = 0,
    Short = 1,
    Int = 2,
    Float = 3,
    String = 4,
    CompoundTag = 5,
    Pos = 6,
    Int64 = 7,
    Vec3 = 8,
}

impl TryFrom<u8> for DataItemType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Byte),
            1 => Ok(Self::Short),
            2 => Ok(Self::Int),
            3 => Ok(Self::Float),
            4 => Ok(Self::String),
            5 => Ok(Self::CompoundTag),
            6 => Ok(Self::Pos),
            7 => Ok(Self::Int64),
            8 => Ok(Self::Vec3),
            value => Err(value),
        }
    }
}

impl From<DataItemType> for u8 {
    fn from(value: DataItemType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum EditorWorldType {
    NonEditor = 0,
    EditorProject = 1,
    EditorTestLevel = 2,
    EditorRealmsUpload = 3,
}

impl TryFrom<i32> for EditorWorldType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::NonEditor),
            1 => Ok(Self::EditorProject),
            2 => Ok(Self::EditorTestLevel),
            3 => Ok(Self::EditorRealmsUpload),
            value => Err(value),
        }
    }
}

impl From<EditorWorldType> for i32 {
    fn from(value: EditorWorldType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u32)]
pub enum EducationEditionOffer {
    None = 0,
    RestOfWorld = 1,
    ChinaDeprecated = 2,
}

impl TryFrom<u32> for EducationEditionOffer {
    type Error = u32;

    fn try_from(value: u32) -> Result<Self, u32> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::RestOfWorld),
            2 => Ok(Self::ChinaDeprecated),
            value => Err(value),
        }
    }
}

impl From<EducationEditionOffer> for u32 {
    fn from(value: EducationEditionOffer) -> Self {
        value as u32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum EnchantType {
    Protection = 0,
    FireProtection = 1,
    FeatherFalling = 2,
    BlastProtection = 3,
    ProjectileProtection = 4,
    Thorns = 5,
    Respiration = 6,
    DepthStrider = 7,
    AquaAffinity = 8,
    Sharpness = 9,
    Smite = 10,
    BaneOfArthropods = 11,
    Knockback = 12,
    FireAspect = 13,
    Looting = 14,
    Efficiency = 15,
    SilkTouch = 16,
    Unbreaking = 17,
    Fortune = 18,
    Power = 19,
    Punch = 20,
    Flame = 21,
    Infinity = 22,
    LuckOfTheSea = 23,
    Lure = 24,
    FrostWalker = 25,
    Mending = 26,
    CurseOfBinding = 27,
    CurseOfVanishing = 28,
    Impaling = 29,
    Riptide = 30,
    Loyalty = 31,
    Channeling = 32,
    Multishot = 33,
    Piercing = 34,
    QuickCharge = 35,
    SoulSpeed = 36,
    SwiftSneak = 37,
    WindBurst = 38,
    Density = 39,
    Breach = 40,
    Lunge = 41,
    NumEnchantments = 42,
    InvalidEnchantment = 43,
}

impl TryFrom<u8> for EnchantType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Protection),
            1 => Ok(Self::FireProtection),
            2 => Ok(Self::FeatherFalling),
            3 => Ok(Self::BlastProtection),
            4 => Ok(Self::ProjectileProtection),
            5 => Ok(Self::Thorns),
            6 => Ok(Self::Respiration),
            7 => Ok(Self::DepthStrider),
            8 => Ok(Self::AquaAffinity),
            9 => Ok(Self::Sharpness),
            10 => Ok(Self::Smite),
            11 => Ok(Self::BaneOfArthropods),
            12 => Ok(Self::Knockback),
            13 => Ok(Self::FireAspect),
            14 => Ok(Self::Looting),
            15 => Ok(Self::Efficiency),
            16 => Ok(Self::SilkTouch),
            17 => Ok(Self::Unbreaking),
            18 => Ok(Self::Fortune),
            19 => Ok(Self::Power),
            20 => Ok(Self::Punch),
            21 => Ok(Self::Flame),
            22 => Ok(Self::Infinity),
            23 => Ok(Self::LuckOfTheSea),
            24 => Ok(Self::Lure),
            25 => Ok(Self::FrostWalker),
            26 => Ok(Self::Mending),
            27 => Ok(Self::CurseOfBinding),
            28 => Ok(Self::CurseOfVanishing),
            29 => Ok(Self::Impaling),
            30 => Ok(Self::Riptide),
            31 => Ok(Self::Loyalty),
            32 => Ok(Self::Channeling),
            33 => Ok(Self::Multishot),
            34 => Ok(Self::Piercing),
            35 => Ok(Self::QuickCharge),
            36 => Ok(Self::SoulSpeed),
            37 => Ok(Self::SwiftSneak),
            38 => Ok(Self::WindBurst),
            39 => Ok(Self::Density),
            40 => Ok(Self::Breach),
            41 => Ok(Self::Lunge),
            42 => Ok(Self::NumEnchantments),
            43 => Ok(Self::InvalidEnchantment),
            value => Err(value),
        }
    }
}

impl From<EnchantType> for u8 {
    fn from(value: EnchantType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum GameType {
    Undefined = -1,
    Survival = 0,
    Creative = 1,
    Adventure = 2,
    Default = 5,
    Spectator = 6,
}

impl TryFrom<i32> for GameType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            -1 => Ok(Self::Undefined),
            0 => Ok(Self::Survival),
            1 => Ok(Self::Creative),
            2 => Ok(Self::Adventure),
            5 => Ok(Self::Default),
            6 => Ok(Self::Spectator),
            value => Err(value),
        }
    }
}

impl From<GameType> for i32 {
    fn from(value: GameType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum GeneratorType {
    Legacy = 0,
    Overworld = 1,
    Flat = 2,
    Nether = 3,
    TheEnd = 4,
    Void = 5,
    Undefined = 6,
}

impl TryFrom<i32> for GeneratorType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Legacy),
            1 => Ok(Self::Overworld),
            2 => Ok(Self::Flat),
            3 => Ok(Self::Nether),
            4 => Ok(Self::TheEnd),
            5 => Ok(Self::Void),
            6 => Ok(Self::Undefined),
            value => Err(value),
        }
    }
}

impl From<GeneratorType> for i32 {
    fn from(value: GeneratorType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum GraphicsMode {
    Simple = 0,
    Fancy = 1,
    Advanced = 2,
    RayTraced = 3,
}

impl TryFrom<u8> for GraphicsMode {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Simple),
            1 => Ok(Self::Fancy),
            2 => Ok(Self::Advanced),
            3 => Ok(Self::RayTraced),
            value => Err(value),
        }
    }
}

impl From<GraphicsMode> for u8 {
    fn from(value: GraphicsMode) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum GraphicsOverrideParameterType {
    SkyZenithColor = 0,
    SkyHorizonColor = 1,
    HorizonBlendMin = 2,
    HorizonBlendMax = 3,
    HorizonBlendStart = 4,
    HorizonBlendMieStart = 5,
    RayleighStrength = 6,
    SunMieStrength = 7,
    MoonMieStrength = 8,
    SunGlareShape = 9,
    Chlorophyll = 10,
    Cdom = 11,
    SuspendedSediment = 12,
    WavesDepth = 13,
    WavesFrequency = 14,
    WavesFrequencyScaling = 15,
    WavesSpeed = 16,
    WavesSpeedScaling = 17,
    WavesShape = 18,
    WavesOctaves = 19,
    WavesMix = 20,
    WavesPull = 21,
    WavesDirectionIncrement = 22,
    MidtonesContrast = 23,
    HighlightsContrast = 24,
    ShadowsContrast = 25,
    HighlightsGain = 26,
    HighlightsGamma = 27,
    HighlightsOffset = 28,
    HighlightsSaturation = 29,
    MidtonesGain = 30,
    MidtonesGamma = 31,
    MidtonesOffset = 32,
    MidtonesSaturation = 33,
    ShadowsGain = 34,
    ShadowsGamma = 35,
    ShadowsOffset = 36,
    ShadowsSaturation = 37,
    HighlightsMin = 38,
    ShadowsMax = 39,
    Temperature = 40,
    SunColor = 41,
    SunIlluminance = 42,
    MoonColor = 43,
    MoonIlluminance = 44,
    FlashColor = 45,
    FlashIlluminance = 46,
    AmbientColor = 47,
    AmbientIlluminance = 48,
    EmissiveDesaturation = 49,
    SkyIntensity = 50,
    OrbitalOffsetDegrees = 51,
}

impl TryFrom<u8> for GraphicsOverrideParameterType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::SkyZenithColor),
            1 => Ok(Self::SkyHorizonColor),
            2 => Ok(Self::HorizonBlendMin),
            3 => Ok(Self::HorizonBlendMax),
            4 => Ok(Self::HorizonBlendStart),
            5 => Ok(Self::HorizonBlendMieStart),
            6 => Ok(Self::RayleighStrength),
            7 => Ok(Self::SunMieStrength),
            8 => Ok(Self::MoonMieStrength),
            9 => Ok(Self::SunGlareShape),
            10 => Ok(Self::Chlorophyll),
            11 => Ok(Self::Cdom),
            12 => Ok(Self::SuspendedSediment),
            13 => Ok(Self::WavesDepth),
            14 => Ok(Self::WavesFrequency),
            15 => Ok(Self::WavesFrequencyScaling),
            16 => Ok(Self::WavesSpeed),
            17 => Ok(Self::WavesSpeedScaling),
            18 => Ok(Self::WavesShape),
            19 => Ok(Self::WavesOctaves),
            20 => Ok(Self::WavesMix),
            21 => Ok(Self::WavesPull),
            22 => Ok(Self::WavesDirectionIncrement),
            23 => Ok(Self::MidtonesContrast),
            24 => Ok(Self::HighlightsContrast),
            25 => Ok(Self::ShadowsContrast),
            26 => Ok(Self::HighlightsGain),
            27 => Ok(Self::HighlightsGamma),
            28 => Ok(Self::HighlightsOffset),
            29 => Ok(Self::HighlightsSaturation),
            30 => Ok(Self::MidtonesGain),
            31 => Ok(Self::MidtonesGamma),
            32 => Ok(Self::MidtonesOffset),
            33 => Ok(Self::MidtonesSaturation),
            34 => Ok(Self::ShadowsGain),
            35 => Ok(Self::ShadowsGamma),
            36 => Ok(Self::ShadowsOffset),
            37 => Ok(Self::ShadowsSaturation),
            38 => Ok(Self::HighlightsMin),
            39 => Ok(Self::ShadowsMax),
            40 => Ok(Self::Temperature),
            41 => Ok(Self::SunColor),
            42 => Ok(Self::SunIlluminance),
            43 => Ok(Self::MoonColor),
            44 => Ok(Self::MoonIlluminance),
            45 => Ok(Self::FlashColor),
            46 => Ok(Self::FlashIlluminance),
            47 => Ok(Self::AmbientColor),
            48 => Ok(Self::AmbientIlluminance),
            49 => Ok(Self::EmissiveDesaturation),
            50 => Ok(Self::SkyIntensity),
            51 => Ok(Self::OrbitalOffsetDegrees),
            value => Err(value),
        }
    }
}

impl From<GraphicsOverrideParameterType> for u8 {
    fn from(value: GraphicsOverrideParameterType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum HudElement {
    PaperDoll = 0,
    Armor = 1,
    ToolTips = 2,
    TouchControls = 3,
    Crosshair = 4,
    HotBar = 5,
    Health = 6,
    ProgressBar = 7,
    Hunger = 8,
    AirBubbles = 9,
    HorseHealth = 10,
    StatusEffects = 11,
    ItemText = 12,
}

impl TryFrom<i32> for HudElement {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::PaperDoll),
            1 => Ok(Self::Armor),
            2 => Ok(Self::ToolTips),
            3 => Ok(Self::TouchControls),
            4 => Ok(Self::Crosshair),
            5 => Ok(Self::HotBar),
            6 => Ok(Self::Health),
            7 => Ok(Self::ProgressBar),
            8 => Ok(Self::Hunger),
            9 => Ok(Self::AirBubbles),
            10 => Ok(Self::HorseHealth),
            11 => Ok(Self::StatusEffects),
            12 => Ok(Self::ItemText),
            value => Err(value),
        }
    }
}

impl From<HudElement> for i32 {
    fn from(value: HudElement) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum HudVisibility {
    Hide = 0,
    Reset = 1,
}

impl TryFrom<i32> for HudVisibility {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Hide),
            1 => Ok(Self::Reset),
            value => Err(value),
        }
    }
}

impl From<HudVisibility> for i32 {
    fn from(value: HudVisibility) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u32)]
pub enum InputMode {
    Undefined = 0,
    Mouse = 1,
    Touch = 2,
    GamePad = 3,
    MotionController = 4,
    Count = 5,
}

impl TryFrom<u32> for InputMode {
    type Error = u32;

    fn try_from(value: u32) -> Result<Self, u32> {
        match value {
            0 => Ok(Self::Undefined),
            1 => Ok(Self::Mouse),
            2 => Ok(Self::Touch),
            3 => Ok(Self::GamePad),
            4 => Ok(Self::MotionController),
            5 => Ok(Self::Count),
            value => Err(value),
        }
    }
}

impl From<InputMode> for u32 {
    fn from(value: InputMode) -> Self {
        value as u32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum InteractAction {
    Invalid = 0,
    StopRiding = 3,
    InteractUpdate = 4,
    NpcOpen = 5,
    OpenInventory = 6,
}

impl TryFrom<u8> for InteractAction {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Invalid),
            3 => Ok(Self::StopRiding),
            4 => Ok(Self::InteractUpdate),
            5 => Ok(Self::NpcOpen),
            6 => Ok(Self::OpenInventory),
            value => Err(value),
        }
    }
}

impl From<InteractAction> for u8 {
    fn from(value: InteractAction) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum InventoryLayout {
    None = 0,
    InventoryOnly = 1,
    Default = 2,
    RecipeBookOnly = 3,
}

impl TryFrom<i32> for InventoryLayout {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::InventoryOnly),
            2 => Ok(Self::Default),
            3 => Ok(Self::RecipeBookOnly),
            value => Err(value),
        }
    }
}

impl From<InventoryLayout> for i32 {
    fn from(value: InventoryLayout) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum InventoryLeftTabIndex {
    None = 0,
    RecipeConstruction = 1,
    RecipeEquipment = 2,
    RecipeItems = 3,
    RecipeNature = 4,
    RecipeSearch = 5,
    Survival = 6,
}

impl TryFrom<i32> for InventoryLeftTabIndex {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::RecipeConstruction),
            2 => Ok(Self::RecipeEquipment),
            3 => Ok(Self::RecipeItems),
            4 => Ok(Self::RecipeNature),
            5 => Ok(Self::RecipeSearch),
            6 => Ok(Self::Survival),
            value => Err(value),
        }
    }
}

impl From<InventoryLeftTabIndex> for i32 {
    fn from(value: InventoryLeftTabIndex) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum InventoryRightTabIndex {
    None = 0,
    FullScreen = 1,
    Crafting = 2,
    Armor = 3,
}

impl TryFrom<i32> for InventoryRightTabIndex {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::FullScreen),
            2 => Ok(Self::Crafting),
            3 => Ok(Self::Armor),
            value => Err(value),
        }
    }
}

impl From<InventoryRightTabIndex> for i32 {
    fn from(value: InventoryRightTabIndex) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u32)]
pub enum InventorySourceInventorySourceFlags {
    NoFlag = 0,
    WorldInteractionRandom = 1,
}

impl TryFrom<u32> for InventorySourceInventorySourceFlags {
    type Error = u32;

    fn try_from(value: u32) -> Result<Self, u32> {
        match value {
            0 => Ok(Self::NoFlag),
            1 => Ok(Self::WorldInteractionRandom),
            value => Err(value),
        }
    }
}

impl From<InventorySourceInventorySourceFlags> for u32 {
    fn from(value: InventorySourceInventorySourceFlags) -> Self {
        value as u32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u32)]
pub enum InventorySourceType {
    ContainerInventory = 0,
    GlobalInventory = 1,
    WorldInteraction = 2,
    CreativeInventory = 3,
    NonImplementedFeatureTodo = 99999,
}

impl TryFrom<u32> for InventorySourceType {
    type Error = u32;

    fn try_from(value: u32) -> Result<Self, u32> {
        match value {
            0 => Ok(Self::ContainerInventory),
            1 => Ok(Self::GlobalInventory),
            2 => Ok(Self::WorldInteraction),
            3 => Ok(Self::CreativeInventory),
            99999 => Ok(Self::NonImplementedFeatureTodo),
            value => Err(value),
        }
    }
}

impl From<InventorySourceType> for u32 {
    fn from(value: InventorySourceType) -> Self {
        value as u32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum ItemReleaseInventoryTransactionActionType {
    Release = 0,
    Use = 1,
}

impl TryFrom<i32> for ItemReleaseInventoryTransactionActionType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Release),
            1 => Ok(Self::Use),
            value => Err(value),
        }
    }
}

impl From<ItemReleaseInventoryTransactionActionType> for i32 {
    fn from(value: ItemReleaseInventoryTransactionActionType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ItemStackNetResult {
    Success = 0,
    Error = 1,
    InvalidRequestActionType = 2,
    ActionRequestNotAllowed = 3,
    ScreenHandlerEndRequestFailed = 4,
    ItemRequestActionHandlerCommitFailed = 5,
    InvalidRequestCraftActionType = 6,
    InvalidCraftRequest = 7,
    InvalidCraftRequestScreen = 8,
    InvalidCraftResult = 9,
    InvalidCraftResultIndex = 10,
    InvalidCraftResultItem = 11,
    InvalidItemNetId = 12,
    MissingCreatedOutputContainer = 13,
    FailedToSetCreatedItemOutputSlot = 14,
    RequestAlreadyInProgress = 15,
    FailedToInitSparseContainer = 16,
    ResultTransferFailed = 17,
    ExpectedItemSlotNotFullyConsumed = 18,
    ExpectedAnywhereItemNotFullyConsumed = 19,
    ItemAlreadyConsumedFromSlot = 20,
    ConsumedTooMuchFromSlot = 21,
    MismatchSlotExpectedConsumedItem = 22,
    MismatchSlotExpectedConsumedItemNetIdVariant = 23,
    FailedToMatchExpectedSlotConsumedItem = 24,
    FailedToMatchExpectedAllowedAnywhereConsumedItem = 25,
    ConsumedItemOutOfAllowedSlotRange = 26,
    ConsumedItemNotAllowed = 27,
    PlayerNotInCreativeMode = 28,
    InvalidExperimentalRecipeRequest = 29,
    FailedToCraftCreative = 30,
    FailedToGetLevelRecipe = 31,
    FailedToFindRecipeByNetId = 32,
    MismatchedCraftingSize = 33,
    MissingInputSparseContainer = 34,
    MismatchedRecipeForInputGridItems = 35,
    EmptyCraftResults = 36,
    FailedToEnchant = 37,
    MissingInputItem = 38,
    InsufficientPlayerLevelToEnchant = 39,
    MissingMaterialItem = 40,
    MissingActor = 41,
    UnknownPrimaryEffect = 42,
    PrimaryEffectOutOfRange = 43,
    PrimaryEffectUnavailable = 44,
    SecondaryEffectOutOfRange = 45,
    SecondaryEffectUnavailable = 46,
    DstContainerEqualToCreatedOutputContainer = 47,
    DstContainerAndSlotEqualToSrcContainerAndSlot = 48,
    FailedToValidateSrcSlot = 49,
    FailedToValidateDstSlot = 50,
    InvalidAdjustedAmount = 51,
    InvalidItemSetType = 52,
    InvalidTransferAmount = 53,
    CannotSwapItem = 54,
    CannotPlaceItem = 55,
    UnhandledItemSetType = 56,
    InvalidRemovedAmount = 57,
    InvalidRegion = 58,
    CannotDropItem = 59,
    CannotDestroyItem = 60,
    InvalidSourceContainer = 61,
    ItemNotConsumed = 62,
    InvalidNumCrafts = 63,
    InvalidCraftResultStackSize = 64,
    CannotRemoveItem = 65,
    CannotConsumeItem = 66,
    ScreenStackError = 67,
}

impl TryFrom<u8> for ItemStackNetResult {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Success),
            1 => Ok(Self::Error),
            2 => Ok(Self::InvalidRequestActionType),
            3 => Ok(Self::ActionRequestNotAllowed),
            4 => Ok(Self::ScreenHandlerEndRequestFailed),
            5 => Ok(Self::ItemRequestActionHandlerCommitFailed),
            6 => Ok(Self::InvalidRequestCraftActionType),
            7 => Ok(Self::InvalidCraftRequest),
            8 => Ok(Self::InvalidCraftRequestScreen),
            9 => Ok(Self::InvalidCraftResult),
            10 => Ok(Self::InvalidCraftResultIndex),
            11 => Ok(Self::InvalidCraftResultItem),
            12 => Ok(Self::InvalidItemNetId),
            13 => Ok(Self::MissingCreatedOutputContainer),
            14 => Ok(Self::FailedToSetCreatedItemOutputSlot),
            15 => Ok(Self::RequestAlreadyInProgress),
            16 => Ok(Self::FailedToInitSparseContainer),
            17 => Ok(Self::ResultTransferFailed),
            18 => Ok(Self::ExpectedItemSlotNotFullyConsumed),
            19 => Ok(Self::ExpectedAnywhereItemNotFullyConsumed),
            20 => Ok(Self::ItemAlreadyConsumedFromSlot),
            21 => Ok(Self::ConsumedTooMuchFromSlot),
            22 => Ok(Self::MismatchSlotExpectedConsumedItem),
            23 => Ok(Self::MismatchSlotExpectedConsumedItemNetIdVariant),
            24 => Ok(Self::FailedToMatchExpectedSlotConsumedItem),
            25 => Ok(Self::FailedToMatchExpectedAllowedAnywhereConsumedItem),
            26 => Ok(Self::ConsumedItemOutOfAllowedSlotRange),
            27 => Ok(Self::ConsumedItemNotAllowed),
            28 => Ok(Self::PlayerNotInCreativeMode),
            29 => Ok(Self::InvalidExperimentalRecipeRequest),
            30 => Ok(Self::FailedToCraftCreative),
            31 => Ok(Self::FailedToGetLevelRecipe),
            32 => Ok(Self::FailedToFindRecipeByNetId),
            33 => Ok(Self::MismatchedCraftingSize),
            34 => Ok(Self::MissingInputSparseContainer),
            35 => Ok(Self::MismatchedRecipeForInputGridItems),
            36 => Ok(Self::EmptyCraftResults),
            37 => Ok(Self::FailedToEnchant),
            38 => Ok(Self::MissingInputItem),
            39 => Ok(Self::InsufficientPlayerLevelToEnchant),
            40 => Ok(Self::MissingMaterialItem),
            41 => Ok(Self::MissingActor),
            42 => Ok(Self::UnknownPrimaryEffect),
            43 => Ok(Self::PrimaryEffectOutOfRange),
            44 => Ok(Self::PrimaryEffectUnavailable),
            45 => Ok(Self::SecondaryEffectOutOfRange),
            46 => Ok(Self::SecondaryEffectUnavailable),
            47 => Ok(Self::DstContainerEqualToCreatedOutputContainer),
            48 => Ok(Self::DstContainerAndSlotEqualToSrcContainerAndSlot),
            49 => Ok(Self::FailedToValidateSrcSlot),
            50 => Ok(Self::FailedToValidateDstSlot),
            51 => Ok(Self::InvalidAdjustedAmount),
            52 => Ok(Self::InvalidItemSetType),
            53 => Ok(Self::InvalidTransferAmount),
            54 => Ok(Self::CannotSwapItem),
            55 => Ok(Self::CannotPlaceItem),
            56 => Ok(Self::UnhandledItemSetType),
            57 => Ok(Self::InvalidRemovedAmount),
            58 => Ok(Self::InvalidRegion),
            59 => Ok(Self::CannotDropItem),
            60 => Ok(Self::CannotDestroyItem),
            61 => Ok(Self::InvalidSourceContainer),
            62 => Ok(Self::ItemNotConsumed),
            63 => Ok(Self::InvalidNumCrafts),
            64 => Ok(Self::InvalidCraftResultStackSize),
            65 => Ok(Self::CannotRemoveItem),
            66 => Ok(Self::CannotConsumeItem),
            67 => Ok(Self::ScreenStackError),
            value => Err(value),
        }
    }
}

impl From<ItemStackNetResult> for u8 {
    fn from(value: ItemStackNetResult) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ItemStackRequestActionType {
    Take = 0,
    Place = 1,
    Swap = 2,
    Drop = 3,
    Destroy = 4,
    Consume = 5,
    Create = 6,
    PlaceInItemContainer = 7,
    TakeFromItemContainer = 8,
    ScreenLabTableCombine = 9,
    ScreenBeaconPayment = 10,
    ScreenHudMineBlock = 11,
    CraftRecipe = 12,
    CraftRecipeAuto = 13,
    CraftCreative = 14,
    CraftRecipeOptional = 15,
    CraftRepairAndDisenchant = 16,
    CraftLoom = 17,
    CraftNonImplemented = 18,
    CraftResults = 19,
}

impl TryFrom<u8> for ItemStackRequestActionType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Take),
            1 => Ok(Self::Place),
            2 => Ok(Self::Swap),
            3 => Ok(Self::Drop),
            4 => Ok(Self::Destroy),
            5 => Ok(Self::Consume),
            6 => Ok(Self::Create),
            7 => Ok(Self::PlaceInItemContainer),
            8 => Ok(Self::TakeFromItemContainer),
            9 => Ok(Self::ScreenLabTableCombine),
            10 => Ok(Self::ScreenBeaconPayment),
            11 => Ok(Self::ScreenHudMineBlock),
            12 => Ok(Self::CraftRecipe),
            13 => Ok(Self::CraftRecipeAuto),
            14 => Ok(Self::CraftCreative),
            15 => Ok(Self::CraftRecipeOptional),
            16 => Ok(Self::CraftRepairAndDisenchant),
            17 => Ok(Self::CraftLoom),
            18 => Ok(Self::CraftNonImplemented),
            19 => Ok(Self::CraftResults),
            value => Err(value),
        }
    }
}

impl From<ItemStackRequestActionType> for u8 {
    fn from(value: ItemStackRequestActionType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ItemStackRequestCerealItemDescriptorType {
    Empty = 0,
    ItemName = 1,
    MoLang = 2,
    ItemTag = 3,
}

impl TryFrom<u8> for ItemStackRequestCerealItemDescriptorType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Empty),
            1 => Ok(Self::ItemName),
            2 => Ok(Self::MoLang),
            3 => Ok(Self::ItemTag),
            value => Err(value),
        }
    }
}

impl From<ItemStackRequestCerealItemDescriptorType> for u8 {
    fn from(value: ItemStackRequestCerealItemDescriptorType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum ItemUseInventoryTransactionActionType {
    Place = 0,
    Use = 1,
    Destroy = 2,
    UseAsAttack = 3,
}

impl TryFrom<i32> for ItemUseInventoryTransactionActionType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Place),
            1 => Ok(Self::Use),
            2 => Ok(Self::Destroy),
            3 => Ok(Self::UseAsAttack),
            value => Err(value),
        }
    }
}

impl From<ItemUseInventoryTransactionActionType> for i32 {
    fn from(value: ItemUseInventoryTransactionActionType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ItemUseInventoryTransactionClientCooldownState {
    Off = 0,
    On = 1,
}

impl TryFrom<u8> for ItemUseInventoryTransactionClientCooldownState {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Off),
            1 => Ok(Self::On),
            value => Err(value),
        }
    }
}

impl From<ItemUseInventoryTransactionClientCooldownState> for u8 {
    fn from(value: ItemUseInventoryTransactionClientCooldownState) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ItemUseInventoryTransactionPredictedResult {
    Failure = 0,
    Success = 1,
}

impl TryFrom<u8> for ItemUseInventoryTransactionPredictedResult {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Failure),
            1 => Ok(Self::Success),
            value => Err(value),
        }
    }
}

impl From<ItemUseInventoryTransactionPredictedResult> for u8 {
    fn from(value: ItemUseInventoryTransactionPredictedResult) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ItemUseInventoryTransactionTriggerType {
    Unknown = 0,
    PlayerInput = 1,
    SimulationTick = 2,
}

impl TryFrom<u8> for ItemUseInventoryTransactionTriggerType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Unknown),
            1 => Ok(Self::PlayerInput),
            2 => Ok(Self::SimulationTick),
            value => Err(value),
        }
    }
}

impl From<ItemUseInventoryTransactionTriggerType> for u8 {
    fn from(value: ItemUseInventoryTransactionTriggerType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum ItemUseOnActorInventoryTransactionActionType {
    Interact = 0,
    Attack = 1,
    ItemInteract = 2,
}

impl TryFrom<i32> for ItemUseOnActorInventoryTransactionActionType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Interact),
            1 => Ok(Self::Attack),
            2 => Ok(Self::ItemInteract),
            value => Err(value),
        }
    }
}

impl From<ItemUseOnActorInventoryTransactionActionType> for i32 {
    fn from(value: ItemUseOnActorInventoryTransactionActionType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum ItemVersion {
    Legacy = 0,
    DataDriven = 1,
    None = 2,
}

impl TryFrom<i32> for ItemVersion {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Legacy),
            1 => Ok(Self::DataDriven),
            2 => Ok(Self::None),
            value => Err(value),
        }
    }
}

impl From<ItemVersion> for i32 {
    fn from(value: ItemVersion) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum LabTableReactionType {
    None = 0,
    IceBomb = 1,
    Bleach = 2,
    ElephantToothpaste = 3,
    Fertilizer = 4,
    HeatBlock = 5,
    MagnesiumSalts = 6,
    MiscFire = 7,
    MiscExplosion = 8,
    MiscLava = 9,
    MiscMystical = 10,
    MiscSmoke = 11,
    MiscLargeSmoke = 12,
}

impl TryFrom<u8> for LabTableReactionType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::IceBomb),
            2 => Ok(Self::Bleach),
            3 => Ok(Self::ElephantToothpaste),
            4 => Ok(Self::Fertilizer),
            5 => Ok(Self::HeatBlock),
            6 => Ok(Self::MagnesiumSalts),
            7 => Ok(Self::MiscFire),
            8 => Ok(Self::MiscExplosion),
            9 => Ok(Self::MiscLava),
            10 => Ok(Self::MiscMystical),
            11 => Ok(Self::MiscSmoke),
            12 => Ok(Self::MiscLargeSmoke),
            value => Err(value),
        }
    }
}

impl From<LabTableReactionType> for u8 {
    fn from(value: LabTableReactionType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum LabTableType {
    StartCombine = 0,
    StartReaction = 1,
    Reset = 2,
}

impl TryFrom<u8> for LabTableType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::StartCombine),
            1 => Ok(Self::StartReaction),
            2 => Ok(Self::Reset),
            value => Err(value),
        }
    }
}

impl From<LabTableType> for u8 {
    fn from(value: LabTableType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum LegacyArmorSlot {
    Head = 0,
    Torso = 1,
    Legs = 2,
    Feet = 3,
    Body = 4,
}

impl TryFrom<i32> for LegacyArmorSlot {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Head),
            1 => Ok(Self::Torso),
            2 => Ok(Self::Legs),
            3 => Ok(Self::Feet),
            4 => Ok(Self::Body),
            value => Err(value),
        }
    }
}

impl From<LegacyArmorSlot> for i32 {
    fn from(value: LegacyArmorSlot) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum LegacyDifficulty {
    Peaceful = 0,
    Easy = 1,
    Normal = 2,
    Hard = 3,
    Count = 4,
    Unknown = 5,
}

impl TryFrom<i32> for LegacyDifficulty {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Peaceful),
            1 => Ok(Self::Easy),
            2 => Ok(Self::Normal),
            3 => Ok(Self::Hard),
            4 => Ok(Self::Count),
            5 => Ok(Self::Unknown),
            value => Err(value),
        }
    }
}

impl From<LegacyDifficulty> for i32 {
    fn from(value: LegacyDifficulty) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum LegacyTelemetryEventType {
    Achievement = 0,
    Interaction = 1,
    PortalCreated = 2,
    PortalUsed = 3,
    MobKilled = 4,
    CauldronUsed = 5,
    PlayerDied = 6,
    BossKilled = 7,
    AgentCommandObsolete = 8,
    AgentCreated = 9,
    PatternRemovedObsolete = 10,
    SlashCommand = 11,
    FishBucketedObsolete = 12,
    MobBorn = 13,
    PetDiedObsolete = 14,
    PoiCauldronUsed = 15,
    ComposterUsed = 16,
    BellUsed = 17,
    ActorDefinition = 18,
    RaidUpdate = 19,
    PlayerMovementAnomalyObsolete = 20,
    PlayerMovementCorrectedObsolete = 21,
    HoneyHarvested = 22,
    TargetBlockHit = 23,
    PiglinBarter = 24,
    PlayerWaxedOrUnwaxedCopper = 25,
    CodeBuilderRuntimeAction = 26,
    CodeBuilderScoreboard = 27,
    StriderRiddenInLavaInOverworld = 28,
    SneakCloseToSculkSensor = 29,
    CarefulRestoration = 30,
    ItemUsed = 31,
}

impl TryFrom<i32> for LegacyTelemetryEventType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Achievement),
            1 => Ok(Self::Interaction),
            2 => Ok(Self::PortalCreated),
            3 => Ok(Self::PortalUsed),
            4 => Ok(Self::MobKilled),
            5 => Ok(Self::CauldronUsed),
            6 => Ok(Self::PlayerDied),
            7 => Ok(Self::BossKilled),
            8 => Ok(Self::AgentCommandObsolete),
            9 => Ok(Self::AgentCreated),
            10 => Ok(Self::PatternRemovedObsolete),
            11 => Ok(Self::SlashCommand),
            12 => Ok(Self::FishBucketedObsolete),
            13 => Ok(Self::MobBorn),
            14 => Ok(Self::PetDiedObsolete),
            15 => Ok(Self::PoiCauldronUsed),
            16 => Ok(Self::ComposterUsed),
            17 => Ok(Self::BellUsed),
            18 => Ok(Self::ActorDefinition),
            19 => Ok(Self::RaidUpdate),
            20 => Ok(Self::PlayerMovementAnomalyObsolete),
            21 => Ok(Self::PlayerMovementCorrectedObsolete),
            22 => Ok(Self::HoneyHarvested),
            23 => Ok(Self::TargetBlockHit),
            24 => Ok(Self::PiglinBarter),
            25 => Ok(Self::PlayerWaxedOrUnwaxedCopper),
            26 => Ok(Self::CodeBuilderRuntimeAction),
            27 => Ok(Self::CodeBuilderScoreboard),
            28 => Ok(Self::StriderRiddenInLavaInOverworld),
            29 => Ok(Self::SneakCloseToSculkSensor),
            30 => Ok(Self::CarefulRestoration),
            31 => Ok(Self::ItemUsed),
            value => Err(value),
        }
    }
}

impl From<LegacyTelemetryEventType> for i32 {
    fn from(value: LegacyTelemetryEventType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i8)]
pub enum MapDecorationType {
    MarkerWhite = 0,
    MarkerGreen = 1,
    MarkerRed = 2,
    MarkerBlue = 3,
    XWhite = 4,
    TriangleRed = 5,
    SquareWhite = 6,
    MarkerSign = 7,
    MarkerPink = 8,
    MarkerOrange = 9,
    MarkerYellow = 10,
    MarkerTeal = 11,
    TriangleGreen = 12,
    SmallSquareWhite = 13,
    Mansion = 14,
    Monument = 15,
    NoDraw = 16,
    VillageDesert = 17,
    VillagePlains = 18,
    VillageSavanna = 19,
    VillageSnowy = 20,
    VillageTaiga = 21,
    JungleTemple = 22,
    WitchHut = 23,
    TrialChambers = 24,
    Count = 25,
}

impl TryFrom<i8> for MapDecorationType {
    type Error = i8;

    fn try_from(value: i8) -> Result<Self, i8> {
        match value {
            0 => Ok(Self::MarkerWhite),
            1 => Ok(Self::MarkerGreen),
            2 => Ok(Self::MarkerRed),
            3 => Ok(Self::MarkerBlue),
            4 => Ok(Self::XWhite),
            5 => Ok(Self::TriangleRed),
            6 => Ok(Self::SquareWhite),
            7 => Ok(Self::MarkerSign),
            8 => Ok(Self::MarkerPink),
            9 => Ok(Self::MarkerOrange),
            10 => Ok(Self::MarkerYellow),
            11 => Ok(Self::MarkerTeal),
            12 => Ok(Self::TriangleGreen),
            13 => Ok(Self::SmallSquareWhite),
            14 => Ok(Self::Mansion),
            15 => Ok(Self::Monument),
            16 => Ok(Self::NoDraw),
            17 => Ok(Self::VillageDesert),
            18 => Ok(Self::VillagePlains),
            19 => Ok(Self::VillageSavanna),
            20 => Ok(Self::VillageSnowy),
            21 => Ok(Self::VillageTaiga),
            22 => Ok(Self::JungleTemple),
            23 => Ok(Self::WitchHut),
            24 => Ok(Self::TrialChambers),
            25 => Ok(Self::Count),
            value => Err(value),
        }
    }
}

impl From<MapDecorationType> for i8 {
    fn from(value: MapDecorationType) -> Self {
        value as i8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum MapItemTrackedActorType {
    Entity = 0,
    BlockEntity = 1,
    Other = 2,
}

impl TryFrom<i32> for MapItemTrackedActorType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Entity),
            1 => Ok(Self::BlockEntity),
            2 => Ok(Self::Other),
            value => Err(value),
        }
    }
}

impl From<MapItemTrackedActorType> for i32 {
    fn from(value: MapItemTrackedActorType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum MemoryMemoryCategory {
    Unknown = 0,
    InvalidSizeUnknown = 1,
    Actor = 2,
    ActorAnimation = 3,
    ActorRendering = 4,
    BlockTickingQueues = 5,
    BiomeStorage = 6,
    Blobs = 7,
    Cereal = 8,
    CircuitSystem = 9,
    Client = 10,
    Commands = 11,
    DbStorage = 12,
    Debug = 13,
    Documentation = 14,
    EcsSystems = 15,
    Fmod = 16,
    Fonts = 17,
    ImGui = 18,
    Input = 19,
    JsonUi = 20,
    JsonUiControlFactoryJson = 21,
    JsonUiControlTree = 22,
    JsonUiControlTreeControlElement = 23,
    JsonUiControlTreePopulateDataBinding = 24,
    JsonUiControlTreePopulateFocus = 25,
    JsonUiControlTreePopulateLayout = 26,
    JsonUiControlTreePopulateOther = 27,
    JsonUiControlTreePopulateSprite = 28,
    JsonUiControlTreePopulateText = 29,
    JsonUiControlTreePopulateTts = 30,
    JsonUiControlTreeVisibility = 31,
    JsonUiCreateUi = 32,
    JsonUiDefs = 33,
    JsonUiLayoutManager = 34,
    JsonUiLayoutManagerRemoveDependencies = 35,
    JsonUiLayoutManagerInitVariable = 36,
    Languages = 37,
    Level = 38,
    LevelStructures = 39,
    LevelChunk = 40,
    LevelChunkGen = 41,
    LevelChunkGenThreadLocal = 42,
    LightVolumeManager = 43,
    Network = 44,
    Marketplace = 45,
    MaterialDragonCompiledDefinition = 46,
    MaterialDragonMaterial = 47,
    MaterialDragonResource = 48,
    MaterialDragonUniformMap = 49,
    MaterialRenderMaterial = 50,
    MaterialRenderMaterialGroup = 51,
    MaterialVariationManager = 52,
    MoLang = 53,
    OreUi = 54,
    OreUiClient = 55,
    PersonaPieces = 56,
    PersonaAnimations = 57,
    PersonaTextures = 58,
    PersonaCharacters = 59,
    PersonaSkinPacks = 60,
    PersonaRepo = 61,
    Player = 62,
    RenderChunk = 63,
    RenderChunkIndexBuffer = 64,
    RenderChunkVertexBuffer = 65,
    Rendering = 66,
    RenderingBgfxInit = 67,
    RenderingBgfxStartFrame = 68,
    RenderingBlockTessellator = 69,
    RenderingEndFrame = 70,
    RenderingGraphicsTasksInit = 71,
    RenderingLibrary = 72,
    RenderingPolygonOperatorPool = 73,
    RenderingPbrTextureData = 74,
    RenderingRenderRegistry = 75,
    RenderingSetup = 76,
    RenderingVertices = 77,
    RequestLog = 78,
    ResourcePacks = 79,
    Sound = 80,
    SubChunkBiomeData = 81,
    SubChunkBlockData = 82,
    SubChunkLightData = 83,
    Textures = 84,
    WeatherRenderer = 85,
    WorldGenerator = 86,
    Tasks = 87,
    Test = 88,
    TestLoadTestTags = 89,
    Scripting = 90,
    ScriptingRuntime = 91,
    ScriptingContext = 92,
    ScriptingContextBindingsMc = 93,
    ScriptingContextBindingsGt = 94,
    ScriptingContextRun = 95,
    DataDrivenUi = 96,
    DataDrivenUiDefs = 97,
    Gameface = 98,
    GamefaceSystem = 99,
    GamefaceDom = 100,
    GamefaceCss = 101,
    GamefaceDisplay = 102,
    GamefaceTempAllocator = 103,
    GamefacePoolAllocator = 104,
    GamefaceDump = 105,
    GamefaceMedia = 106,
    GamefaceJson = 107,
    GamefaceScriptEngine = 108,
    GamefaceScript = 109,
    GamefaceLayout = 110,
}

impl TryFrom<u8> for MemoryMemoryCategory {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Unknown),
            1 => Ok(Self::InvalidSizeUnknown),
            2 => Ok(Self::Actor),
            3 => Ok(Self::ActorAnimation),
            4 => Ok(Self::ActorRendering),
            5 => Ok(Self::BlockTickingQueues),
            6 => Ok(Self::BiomeStorage),
            7 => Ok(Self::Blobs),
            8 => Ok(Self::Cereal),
            9 => Ok(Self::CircuitSystem),
            10 => Ok(Self::Client),
            11 => Ok(Self::Commands),
            12 => Ok(Self::DbStorage),
            13 => Ok(Self::Debug),
            14 => Ok(Self::Documentation),
            15 => Ok(Self::EcsSystems),
            16 => Ok(Self::Fmod),
            17 => Ok(Self::Fonts),
            18 => Ok(Self::ImGui),
            19 => Ok(Self::Input),
            20 => Ok(Self::JsonUi),
            21 => Ok(Self::JsonUiControlFactoryJson),
            22 => Ok(Self::JsonUiControlTree),
            23 => Ok(Self::JsonUiControlTreeControlElement),
            24 => Ok(Self::JsonUiControlTreePopulateDataBinding),
            25 => Ok(Self::JsonUiControlTreePopulateFocus),
            26 => Ok(Self::JsonUiControlTreePopulateLayout),
            27 => Ok(Self::JsonUiControlTreePopulateOther),
            28 => Ok(Self::JsonUiControlTreePopulateSprite),
            29 => Ok(Self::JsonUiControlTreePopulateText),
            30 => Ok(Self::JsonUiControlTreePopulateTts),
            31 => Ok(Self::JsonUiControlTreeVisibility),
            32 => Ok(Self::JsonUiCreateUi),
            33 => Ok(Self::JsonUiDefs),
            34 => Ok(Self::JsonUiLayoutManager),
            35 => Ok(Self::JsonUiLayoutManagerRemoveDependencies),
            36 => Ok(Self::JsonUiLayoutManagerInitVariable),
            37 => Ok(Self::Languages),
            38 => Ok(Self::Level),
            39 => Ok(Self::LevelStructures),
            40 => Ok(Self::LevelChunk),
            41 => Ok(Self::LevelChunkGen),
            42 => Ok(Self::LevelChunkGenThreadLocal),
            43 => Ok(Self::LightVolumeManager),
            44 => Ok(Self::Network),
            45 => Ok(Self::Marketplace),
            46 => Ok(Self::MaterialDragonCompiledDefinition),
            47 => Ok(Self::MaterialDragonMaterial),
            48 => Ok(Self::MaterialDragonResource),
            49 => Ok(Self::MaterialDragonUniformMap),
            50 => Ok(Self::MaterialRenderMaterial),
            51 => Ok(Self::MaterialRenderMaterialGroup),
            52 => Ok(Self::MaterialVariationManager),
            53 => Ok(Self::MoLang),
            54 => Ok(Self::OreUi),
            55 => Ok(Self::OreUiClient),
            56 => Ok(Self::PersonaPieces),
            57 => Ok(Self::PersonaAnimations),
            58 => Ok(Self::PersonaTextures),
            59 => Ok(Self::PersonaCharacters),
            60 => Ok(Self::PersonaSkinPacks),
            61 => Ok(Self::PersonaRepo),
            62 => Ok(Self::Player),
            63 => Ok(Self::RenderChunk),
            64 => Ok(Self::RenderChunkIndexBuffer),
            65 => Ok(Self::RenderChunkVertexBuffer),
            66 => Ok(Self::Rendering),
            67 => Ok(Self::RenderingBgfxInit),
            68 => Ok(Self::RenderingBgfxStartFrame),
            69 => Ok(Self::RenderingBlockTessellator),
            70 => Ok(Self::RenderingEndFrame),
            71 => Ok(Self::RenderingGraphicsTasksInit),
            72 => Ok(Self::RenderingLibrary),
            73 => Ok(Self::RenderingPolygonOperatorPool),
            74 => Ok(Self::RenderingPbrTextureData),
            75 => Ok(Self::RenderingRenderRegistry),
            76 => Ok(Self::RenderingSetup),
            77 => Ok(Self::RenderingVertices),
            78 => Ok(Self::RequestLog),
            79 => Ok(Self::ResourcePacks),
            80 => Ok(Self::Sound),
            81 => Ok(Self::SubChunkBiomeData),
            82 => Ok(Self::SubChunkBlockData),
            83 => Ok(Self::SubChunkLightData),
            84 => Ok(Self::Textures),
            85 => Ok(Self::WeatherRenderer),
            86 => Ok(Self::WorldGenerator),
            87 => Ok(Self::Tasks),
            88 => Ok(Self::Test),
            89 => Ok(Self::TestLoadTestTags),
            90 => Ok(Self::Scripting),
            91 => Ok(Self::ScriptingRuntime),
            92 => Ok(Self::ScriptingContext),
            93 => Ok(Self::ScriptingContextBindingsMc),
            94 => Ok(Self::ScriptingContextBindingsGt),
            95 => Ok(Self::ScriptingContextRun),
            96 => Ok(Self::DataDrivenUi),
            97 => Ok(Self::DataDrivenUiDefs),
            98 => Ok(Self::Gameface),
            99 => Ok(Self::GamefaceSystem),
            100 => Ok(Self::GamefaceDom),
            101 => Ok(Self::GamefaceCss),
            102 => Ok(Self::GamefaceDisplay),
            103 => Ok(Self::GamefaceTempAllocator),
            104 => Ok(Self::GamefacePoolAllocator),
            105 => Ok(Self::GamefaceDump),
            106 => Ok(Self::GamefaceMedia),
            107 => Ok(Self::GamefaceJson),
            108 => Ok(Self::GamefaceScriptEngine),
            109 => Ok(Self::GamefaceScript),
            110 => Ok(Self::GamefaceLayout),
            value => Err(value),
        }
    }
}

impl From<MemoryMemoryCategory> for u8 {
    fn from(value: MemoryMemoryCategory) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum MinecraftEventingAchievementIds {
    ChestFullOfCobblestone = 7,
    DiamondForYou = 10,
    IronBelly = 20,
    IronMan = 21,
    OnARail = 29,
    Overkill = 30,
    ReturnToSender = 37,
    SniperDuel = 38,
    StayinFrosty = 39,
    TakeInventory = 40,
    MapRoom = 50,
    FreightStation = 52,
    SmeltEverything = 53,
    TasteOfYourOwnMedicine = 54,
    WhenPigsFly = 56,
    Inception = 58,
    ArtificialSelection = 60,
    FreeDiver = 61,
    SpawnTheWither = 62,
    Beaconator = 63,
    GreatView = 64,
    SuperSonic = 65,
    TheEndAgain = 66,
    TreasureHunter = 67,
    ShootingStar = 68,
    FashionShow = 69,
    SelfPublishedAuthor = 71,
    AlternativeFuel = 72,
    SleepWithTheFishes = 73,
    Castaway = 74,
    ImAMarineBiologist = 75,
    SailThe7Seas = 76,
    MeGold = 77,
    Ahoy = 78,
    Atlantis = 79,
    OnePickleTwoPickleSeaPickleFour = 80,
    DoaBarrelRoll = 81,
    Moskstraumen = 82,
    Echolocation = 83,
    WhereHaveYouBeen = 84,
    TopOfTheWorld = 85,
    FruitOnTheLoom = 86,
    SoundTheAlarm = 87,
    BuyLowSellHigh = 88,
    Disenchanted = 89,
    TimeForStew = 90,
    BeeOurGuest = 91,
    TotalBeeLocation = 92,
    StickySituation = 93,
    CoverMeInDebris = 94,
    FloatYourGoat = 95,
    Friend = 96,
    WaxOnWaxOff = 97,
    StriderRiddenInLavaInOverworld = 98,
    GoatHornAcquired = 99,
    JukeboxUsedInMeadows = 100,
    TradedAtWorldHeight = 101,
    SurvivedFallFromWorldHeight = 102,
    SneakCloseToSculkSensor = 103,
    ItSpreads = 104,
    BirthdaySong = 105,
    WithOurPowersCombined = 106,
    PlantingThePast = 107,
    CarefulRestoration = 108,
    Revaulting = 109,
    CraftersCraftingCrafters = 110,
    WhoNeedsRockets = 111,
    OverOverkill = 112,
    HeartTransplanter = 113,
    StayHydrated = 114,
    MobKabob = 115,
    AdventuringTime = 116,
    UhOh = 117,
    GettingWood = 118,
    BenchMaking = 119,
    TimeToMine = 120,
    HotTopic = 121,
    AcquireHardware = 122,
    GettingAnUpgrade = 123,
    MonsterHunter = 124,
    Diamonds = 125,
    PlethoraOfCats = 126,
}

impl TryFrom<u8> for MinecraftEventingAchievementIds {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            7 => Ok(Self::ChestFullOfCobblestone),
            10 => Ok(Self::DiamondForYou),
            20 => Ok(Self::IronBelly),
            21 => Ok(Self::IronMan),
            29 => Ok(Self::OnARail),
            30 => Ok(Self::Overkill),
            37 => Ok(Self::ReturnToSender),
            38 => Ok(Self::SniperDuel),
            39 => Ok(Self::StayinFrosty),
            40 => Ok(Self::TakeInventory),
            50 => Ok(Self::MapRoom),
            52 => Ok(Self::FreightStation),
            53 => Ok(Self::SmeltEverything),
            54 => Ok(Self::TasteOfYourOwnMedicine),
            56 => Ok(Self::WhenPigsFly),
            58 => Ok(Self::Inception),
            60 => Ok(Self::ArtificialSelection),
            61 => Ok(Self::FreeDiver),
            62 => Ok(Self::SpawnTheWither),
            63 => Ok(Self::Beaconator),
            64 => Ok(Self::GreatView),
            65 => Ok(Self::SuperSonic),
            66 => Ok(Self::TheEndAgain),
            67 => Ok(Self::TreasureHunter),
            68 => Ok(Self::ShootingStar),
            69 => Ok(Self::FashionShow),
            71 => Ok(Self::SelfPublishedAuthor),
            72 => Ok(Self::AlternativeFuel),
            73 => Ok(Self::SleepWithTheFishes),
            74 => Ok(Self::Castaway),
            75 => Ok(Self::ImAMarineBiologist),
            76 => Ok(Self::SailThe7Seas),
            77 => Ok(Self::MeGold),
            78 => Ok(Self::Ahoy),
            79 => Ok(Self::Atlantis),
            80 => Ok(Self::OnePickleTwoPickleSeaPickleFour),
            81 => Ok(Self::DoaBarrelRoll),
            82 => Ok(Self::Moskstraumen),
            83 => Ok(Self::Echolocation),
            84 => Ok(Self::WhereHaveYouBeen),
            85 => Ok(Self::TopOfTheWorld),
            86 => Ok(Self::FruitOnTheLoom),
            87 => Ok(Self::SoundTheAlarm),
            88 => Ok(Self::BuyLowSellHigh),
            89 => Ok(Self::Disenchanted),
            90 => Ok(Self::TimeForStew),
            91 => Ok(Self::BeeOurGuest),
            92 => Ok(Self::TotalBeeLocation),
            93 => Ok(Self::StickySituation),
            94 => Ok(Self::CoverMeInDebris),
            95 => Ok(Self::FloatYourGoat),
            96 => Ok(Self::Friend),
            97 => Ok(Self::WaxOnWaxOff),
            98 => Ok(Self::StriderRiddenInLavaInOverworld),
            99 => Ok(Self::GoatHornAcquired),
            100 => Ok(Self::JukeboxUsedInMeadows),
            101 => Ok(Self::TradedAtWorldHeight),
            102 => Ok(Self::SurvivedFallFromWorldHeight),
            103 => Ok(Self::SneakCloseToSculkSensor),
            104 => Ok(Self::ItSpreads),
            105 => Ok(Self::BirthdaySong),
            106 => Ok(Self::WithOurPowersCombined),
            107 => Ok(Self::PlantingThePast),
            108 => Ok(Self::CarefulRestoration),
            109 => Ok(Self::Revaulting),
            110 => Ok(Self::CraftersCraftingCrafters),
            111 => Ok(Self::WhoNeedsRockets),
            112 => Ok(Self::OverOverkill),
            113 => Ok(Self::HeartTransplanter),
            114 => Ok(Self::StayHydrated),
            115 => Ok(Self::MobKabob),
            116 => Ok(Self::AdventuringTime),
            117 => Ok(Self::UhOh),
            118 => Ok(Self::GettingWood),
            119 => Ok(Self::BenchMaking),
            120 => Ok(Self::TimeToMine),
            121 => Ok(Self::HotTopic),
            122 => Ok(Self::AcquireHardware),
            123 => Ok(Self::GettingAnUpgrade),
            124 => Ok(Self::MonsterHunter),
            125 => Ok(Self::Diamonds),
            126 => Ok(Self::PlethoraOfCats),
            value => Err(value),
        }
    }
}

impl From<MinecraftEventingAchievementIds> for u8 {
    fn from(value: MinecraftEventingAchievementIds) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum MinecraftEventingInteractionType {
    Breeding = 1,
    Taming = 2,
    Curing = 3,
    Crafted = 4,
    Shearing = 5,
    Milking = 6,
    Trading = 7,
    Feeding = 8,
    Igniting = 9,
    Coloring = 10,
    Naming = 11,
    Leashing = 12,
    Unleashing = 13,
    PetSleep = 14,
    Trusting = 15,
    Commanding = 16,
    Equipping = 17,
}

impl TryFrom<u8> for MinecraftEventingInteractionType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            1 => Ok(Self::Breeding),
            2 => Ok(Self::Taming),
            3 => Ok(Self::Curing),
            4 => Ok(Self::Crafted),
            5 => Ok(Self::Shearing),
            6 => Ok(Self::Milking),
            7 => Ok(Self::Trading),
            8 => Ok(Self::Feeding),
            9 => Ok(Self::Igniting),
            10 => Ok(Self::Coloring),
            11 => Ok(Self::Naming),
            12 => Ok(Self::Leashing),
            13 => Ok(Self::Unleashing),
            14 => Ok(Self::PetSleep),
            15 => Ok(Self::Trusting),
            16 => Ok(Self::Commanding),
            17 => Ok(Self::Equipping),
            value => Err(value),
        }
    }
}

impl From<MinecraftEventingInteractionType> for u8 {
    fn from(value: MinecraftEventingInteractionType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum MinecraftEventingPOIBlockInteractionType {
    None = 0,
    Extend = 1,
    Clone = 2,
    Lock = 3,
    Create = 4,
    CreateLocator = 5,
    Rename = 6,
    ItemPlaced = 7,
    ItemRemoved = 8,
    Cooking = 9,
    Dousing = 10,
    Lighting = 11,
    Haystack = 12,
    Filled = 13,
    Emptied = 14,
    AddDye = 15,
    DyeItem = 16,
    ClearItem = 17,
    EnchantArrow = 18,
    CompostItemPlaced = 19,
    RecoveredBonemeal = 20,
    BookPlaced = 21,
    BookOpened = 22,
    Disenchant = 23,
    Repair = 24,
    DisenchantAndRepair = 25,
}

impl TryFrom<u8> for MinecraftEventingPOIBlockInteractionType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::Extend),
            2 => Ok(Self::Clone),
            3 => Ok(Self::Lock),
            4 => Ok(Self::Create),
            5 => Ok(Self::CreateLocator),
            6 => Ok(Self::Rename),
            7 => Ok(Self::ItemPlaced),
            8 => Ok(Self::ItemRemoved),
            9 => Ok(Self::Cooking),
            10 => Ok(Self::Dousing),
            11 => Ok(Self::Lighting),
            12 => Ok(Self::Haystack),
            13 => Ok(Self::Filled),
            14 => Ok(Self::Emptied),
            15 => Ok(Self::AddDye),
            16 => Ok(Self::DyeItem),
            17 => Ok(Self::ClearItem),
            18 => Ok(Self::EnchantArrow),
            19 => Ok(Self::CompostItemPlaced),
            20 => Ok(Self::RecoveredBonemeal),
            21 => Ok(Self::BookPlaced),
            22 => Ok(Self::BookOpened),
            23 => Ok(Self::Disenchant),
            24 => Ok(Self::Repair),
            25 => Ok(Self::DisenchantAndRepair),
            value => Err(value),
        }
    }
}

impl From<MinecraftEventingPOIBlockInteractionType> for u8 {
    fn from(value: MinecraftEventingPOIBlockInteractionType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum Mirror {
    None = 0,
    X = 1,
    Z = 2,
    Xz = 3,
}

impl TryFrom<u8> for Mirror {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::X),
            2 => Ok(Self::Z),
            3 => Ok(Self::Xz),
            value => Err(value),
        }
    }
}

impl From<Mirror> for u8 {
    fn from(value: Mirror) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i16)]
pub enum MoLangVersion {
    Invalid = -1,
    BeforeVersioning = 0,
    Initial = 1,
    FixedItemRemainingUseDurationQuery = 2,
    ExpressionErrorMessages = 3,
    UnexpectedOperatorErrors = 4,
    ConditionalOperatorAssociativity = 5,
    ComparisonAndLogicalOperatorPrecedence = 6,
    DivideByNegativeValue = 7,
    FixedCapeFlapAmountQuery = 8,
    QueryBlockPropertyRenamedToState = 9,
    DeprecateOldBlockQueryNames = 10,
    DeprecatedSnifferAndCamelQueries = 11,
    LeafSupportingInFirstSolidBlockBelow = 12,
    Latest = 13,
    NumValidVersions = 14,
}

impl TryFrom<i16> for MoLangVersion {
    type Error = i16;

    fn try_from(value: i16) -> Result<Self, i16> {
        match value {
            -1 => Ok(Self::Invalid),
            0 => Ok(Self::BeforeVersioning),
            1 => Ok(Self::Initial),
            2 => Ok(Self::FixedItemRemainingUseDurationQuery),
            3 => Ok(Self::ExpressionErrorMessages),
            4 => Ok(Self::UnexpectedOperatorErrors),
            5 => Ok(Self::ConditionalOperatorAssociativity),
            6 => Ok(Self::ComparisonAndLogicalOperatorPrecedence),
            7 => Ok(Self::DivideByNegativeValue),
            8 => Ok(Self::FixedCapeFlapAmountQuery),
            9 => Ok(Self::QueryBlockPropertyRenamedToState),
            10 => Ok(Self::DeprecateOldBlockQueryNames),
            11 => Ok(Self::DeprecatedSnifferAndCamelQueries),
            12 => Ok(Self::LeafSupportingInFirstSolidBlockBelow),
            13 => Ok(Self::Latest),
            14 => Ok(Self::NumValidVersions),
            value => Err(value),
        }
    }
}

impl From<MoLangVersion> for i16 {
    fn from(value: MoLangVersion) -> Self {
        value as i16
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum MobEffectEvent {
    Invalid = 0,
    Add = 1,
    Update = 2,
    Remove = 3,
}

impl TryFrom<u8> for MobEffectEvent {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Invalid),
            1 => Ok(Self::Add),
            2 => Ok(Self::Update),
            3 => Ok(Self::Remove),
            value => Err(value),
        }
    }
}

impl From<MobEffectEvent> for u8 {
    fn from(value: MobEffectEvent) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ModalFormCancelReason {
    UserClosed = 0,
    UserBusy = 1,
}

impl TryFrom<u8> for ModalFormCancelReason {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::UserClosed),
            1 => Ok(Self::UserBusy),
            value => Err(value),
        }
    }
}

impl From<ModalFormCancelReason> for u8 {
    fn from(value: ModalFormCancelReason) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum MovementEffectType {
    GlideBoost = 0,
    DolphinBoost = 1,
    GeyserBoost = 2,
}

impl TryFrom<i32> for MovementEffectType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::GlideBoost),
            1 => Ok(Self::DolphinBoost),
            2 => Ok(Self::GeyserBoost),
            value => Err(value),
        }
    }
}

impl From<MovementEffectType> for i32 {
    fn from(value: MovementEffectType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum MultiplayerSettingsPacketType {
    Enable = 0,
    Disable = 1,
    RefreshJoinCode = 2,
}

impl TryFrom<i32> for MultiplayerSettingsPacketType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Enable),
            1 => Ok(Self::Disable),
            2 => Ok(Self::RefreshJoinCode),
            value => Err(value),
        }
    }
}

impl From<MultiplayerSettingsPacketType> for i32 {
    fn from(value: MultiplayerSettingsPacketType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum NewInteractionModel {
    Touch = 0,
    Crosshair = 1,
    Classic = 2,
    Count = 3,
}

impl TryFrom<i32> for NewInteractionModel {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Touch),
            1 => Ok(Self::Crosshair),
            2 => Ok(Self::Classic),
            3 => Ok(Self::Count),
            value => Err(value),
        }
    }
}

impl From<NewInteractionModel> for i32 {
    fn from(value: NewInteractionModel) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum NpcDialogueNpcDialogueActionType {
    Open = 0,
    Close = 1,
}

impl TryFrom<i32> for NpcDialogueNpcDialogueActionType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Open),
            1 => Ok(Self::Close),
            value => Err(value),
        }
    }
}

impl From<NpcDialogueNpcDialogueActionType> for i32 {
    fn from(value: NpcDialogueNpcDialogueActionType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum NpcRequestRequestType {
    SetActions = 0,
    ExecuteAction = 1,
    ExecuteClosingCommands = 2,
    SetName = 3,
    SetSkin = 4,
    SetInteractText = 5,
    ExecuteOpeningCommands = 6,
}

impl TryFrom<u8> for NpcRequestRequestType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::SetActions),
            1 => Ok(Self::ExecuteAction),
            2 => Ok(Self::ExecuteClosingCommands),
            3 => Ok(Self::SetName),
            4 => Ok(Self::SetSkin),
            5 => Ok(Self::SetInteractText),
            6 => Ok(Self::ExecuteOpeningCommands),
            value => Err(value),
        }
    }
}

impl From<NpcRequestRequestType> for u8 {
    fn from(value: NpcRequestRequestType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u16)]
pub enum PacketCompressionAlgorithm {
    ZLib = 0,
    Snappy = 1,
    None = 65535,
}

impl TryFrom<u16> for PacketCompressionAlgorithm {
    type Error = u16;

    fn try_from(value: u16) -> Result<Self, u16> {
        match value {
            0 => Ok(Self::ZLib),
            1 => Ok(Self::Snappy),
            65535 => Ok(Self::None),
            value => Err(value),
        }
    }
}

impl From<PacketCompressionAlgorithm> for u16 {
    fn from(value: PacketCompressionAlgorithm) -> Self {
        value as u16
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum PacketViolationSeverity {
    Unknown = -1,
    Warning = 0,
    FinalWarning = 1,
    TerminatingConnection = 2,
}

impl TryFrom<i32> for PacketViolationSeverity {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            -1 => Ok(Self::Unknown),
            0 => Ok(Self::Warning),
            1 => Ok(Self::FinalWarning),
            2 => Ok(Self::TerminatingConnection),
            value => Err(value),
        }
    }
}

impl From<PacketViolationSeverity> for i32 {
    fn from(value: PacketViolationSeverity) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum PacketViolationType {
    Unknown = -1,
    PacketMalformed = 0,
}

impl TryFrom<i32> for PacketViolationType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            -1 => Ok(Self::Unknown),
            0 => Ok(Self::PacketMalformed),
            value => Err(value),
        }
    }
}

impl From<PacketViolationType> for i32 {
    fn from(value: PacketViolationType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u32)]
pub enum PersonaAnimatedTextureType {
    Face = 1,
    Body32x32 = 2,
    Body128x128 = 3,
}

impl TryFrom<u32> for PersonaAnimatedTextureType {
    type Error = u32;

    fn try_from(value: u32) -> Result<Self, u32> {
        match value {
            1 => Ok(Self::Face),
            2 => Ok(Self::Body32x32),
            3 => Ok(Self::Body128x128),
            value => Err(value),
        }
    }
}

impl From<PersonaAnimatedTextureType> for u32 {
    fn from(value: PersonaAnimatedTextureType) -> Self {
        value as u32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u32)]
pub enum PersonaAnimationExpression {
    Linear = 0,
    Blinking = 1,
}

impl TryFrom<u32> for PersonaAnimationExpression {
    type Error = u32;

    fn try_from(value: u32) -> Result<Self, u32> {
        match value {
            0 => Ok(Self::Linear),
            1 => Ok(Self::Blinking),
            value => Err(value),
        }
    }
}

impl From<PersonaAnimationExpression> for u32 {
    fn from(value: PersonaAnimationExpression) -> Self {
        value as u32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum PersonaArmSizeType {
    Slim = 0,
    Wide = 1,
}

impl TryFrom<u8> for PersonaArmSizeType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Slim),
            1 => Ok(Self::Wide),
            value => Err(value),
        }
    }
}

impl From<PersonaArmSizeType> for u8 {
    fn from(value: PersonaArmSizeType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u32)]
pub enum PersonaPieceType {
    Skeleton = 1,
    Body = 2,
    Skin = 3,
    Bottom = 4,
    Feet = 5,
    Dress = 6,
    Top = 7,
    HighPants = 8,
    Hands = 9,
    Outerwear = 10,
    FacialHair = 11,
    Mouth = 12,
    Eyes = 13,
    Hair = 14,
    Hood = 15,
    Back = 16,
    FaceAccessory = 17,
    Head = 18,
    Legs = 19,
    LeftLeg = 20,
    RightLeg = 21,
    Arms = 22,
    LeftArm = 23,
    RightArm = 24,
    Capes = 25,
    ClassicSkin = 26,
    Emote = 27,
}

impl TryFrom<u32> for PersonaPieceType {
    type Error = u32;

    fn try_from(value: u32) -> Result<Self, u32> {
        match value {
            1 => Ok(Self::Skeleton),
            2 => Ok(Self::Body),
            3 => Ok(Self::Skin),
            4 => Ok(Self::Bottom),
            5 => Ok(Self::Feet),
            6 => Ok(Self::Dress),
            7 => Ok(Self::Top),
            8 => Ok(Self::HighPants),
            9 => Ok(Self::Hands),
            10 => Ok(Self::Outerwear),
            11 => Ok(Self::FacialHair),
            12 => Ok(Self::Mouth),
            13 => Ok(Self::Eyes),
            14 => Ok(Self::Hair),
            15 => Ok(Self::Hood),
            16 => Ok(Self::Back),
            17 => Ok(Self::FaceAccessory),
            18 => Ok(Self::Head),
            19 => Ok(Self::Legs),
            20 => Ok(Self::LeftLeg),
            21 => Ok(Self::RightLeg),
            22 => Ok(Self::Arms),
            23 => Ok(Self::LeftArm),
            24 => Ok(Self::RightArm),
            25 => Ok(Self::Capes),
            26 => Ok(Self::ClassicSkin),
            27 => Ok(Self::Emote),
            value => Err(value),
        }
    }
}

impl From<PersonaPieceType> for u32 {
    fn from(value: PersonaPieceType) -> Self {
        value as u32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum PhotoType {
    Portfolio = 0,
    PhotoItem = 1,
    Book = 2,
}

impl TryFrom<u8> for PhotoType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Portfolio),
            1 => Ok(Self::PhotoItem),
            2 => Ok(Self::Book),
            value => Err(value),
        }
    }
}

impl From<PhotoType> for u8 {
    fn from(value: PhotoType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum PlayStatusType {
    LoginSuccess = 0,
    LoginFailedClientOld = 1,
    LoginFailedServerOld = 2,
    PlayerSpawn = 3,
    LoginFailedInvalidTenant = 4,
    LoginFailedEditionMismatchEduToVanilla = 5,
    LoginFailedEditionMismatchVanillaToEdu = 6,
    LoginFailedServerFullSubClient = 7,
    LoginFailedEditorMismatchEditorToVanilla = 8,
    LoginFailedEditorMismatchVanillaToEditor = 9,
}

impl TryFrom<i32> for PlayStatusType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::LoginSuccess),
            1 => Ok(Self::LoginFailedClientOld),
            2 => Ok(Self::LoginFailedServerOld),
            3 => Ok(Self::PlayerSpawn),
            4 => Ok(Self::LoginFailedInvalidTenant),
            5 => Ok(Self::LoginFailedEditionMismatchEduToVanilla),
            6 => Ok(Self::LoginFailedEditionMismatchVanillaToEdu),
            7 => Ok(Self::LoginFailedServerFullSubClient),
            8 => Ok(Self::LoginFailedEditorMismatchEditorToVanilla),
            9 => Ok(Self::LoginFailedEditorMismatchVanillaToEditor),
            value => Err(value),
        }
    }
}

impl From<PlayStatusType> for i32 {
    fn from(value: PlayStatusType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum PlayerActionType {
    Unknown = -1,
    StartDestroyBlock = 0,
    AbortDestroyBlock = 1,
    StopDestroyBlock = 2,
    GetUpdatedBlock = 3,
    DropItem = 4,
    StartSleeping = 5,
    StopSleeping = 6,
    Respawn = 7,
    StartJump = 8,
    StartSprinting = 9,
    StopSprinting = 10,
    StartSneaking = 11,
    StopSneaking = 12,
    CreativeDestroyBlock = 13,
    ChangeDimensionAck = 14,
    StartGliding = 15,
    StopGliding = 16,
    DenyDestroyBlock = 17,
    CrackBlock = 18,
    ChangeSkin = 19,
    UpdatedEnchantingSeed = 20,
    StartSwimming = 21,
    StopSwimming = 22,
    StartSpinAttack = 23,
    StopSpinAttack = 24,
    InteractWithBlock = 25,
    PredictDestroyBlock = 26,
    ContinueDestroyBlock = 27,
    StartItemUseOn = 28,
    StopItemUseOn = 29,
    HandledTeleport = 30,
    MissedSwing = 31,
    StartCrawling = 32,
    StopCrawling = 33,
    StartFlying = 34,
    StopFlying = 35,
    ClientAckServerData = 36,
    StartUsingItem = 37,
    InternalUpdate = 38,
    Count = 39,
}

impl TryFrom<i32> for PlayerActionType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            -1 => Ok(Self::Unknown),
            0 => Ok(Self::StartDestroyBlock),
            1 => Ok(Self::AbortDestroyBlock),
            2 => Ok(Self::StopDestroyBlock),
            3 => Ok(Self::GetUpdatedBlock),
            4 => Ok(Self::DropItem),
            5 => Ok(Self::StartSleeping),
            6 => Ok(Self::StopSleeping),
            7 => Ok(Self::Respawn),
            8 => Ok(Self::StartJump),
            9 => Ok(Self::StartSprinting),
            10 => Ok(Self::StopSprinting),
            11 => Ok(Self::StartSneaking),
            12 => Ok(Self::StopSneaking),
            13 => Ok(Self::CreativeDestroyBlock),
            14 => Ok(Self::ChangeDimensionAck),
            15 => Ok(Self::StartGliding),
            16 => Ok(Self::StopGliding),
            17 => Ok(Self::DenyDestroyBlock),
            18 => Ok(Self::CrackBlock),
            19 => Ok(Self::ChangeSkin),
            20 => Ok(Self::UpdatedEnchantingSeed),
            21 => Ok(Self::StartSwimming),
            22 => Ok(Self::StopSwimming),
            23 => Ok(Self::StartSpinAttack),
            24 => Ok(Self::StopSpinAttack),
            25 => Ok(Self::InteractWithBlock),
            26 => Ok(Self::PredictDestroyBlock),
            27 => Ok(Self::ContinueDestroyBlock),
            28 => Ok(Self::StartItemUseOn),
            29 => Ok(Self::StopItemUseOn),
            30 => Ok(Self::HandledTeleport),
            31 => Ok(Self::MissedSwing),
            32 => Ok(Self::StartCrawling),
            33 => Ok(Self::StopCrawling),
            34 => Ok(Self::StartFlying),
            35 => Ok(Self::StopFlying),
            36 => Ok(Self::ClientAckServerData),
            37 => Ok(Self::StartUsingItem),
            38 => Ok(Self::InternalUpdate),
            39 => Ok(Self::Count),
            value => Err(value),
        }
    }
}

impl From<PlayerActionType> for i32 {
    fn from(value: PlayerActionType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum PlayerAuthInputInputData {
    Ascend = 0,
    Descend = 1,
    NorthJump = 2,
    JumpDown = 3,
    SprintDown = 4,
    ChangeHeight = 5,
    Jumping = 6,
    AutoJumpingInWater = 7,
    Sneaking = 8,
    SneakDown = 9,
    Up = 10,
    Down = 11,
    Left = 12,
    Right = 13,
    UpLeft = 14,
    UpRight = 15,
    WantUp = 16,
    WantDown = 17,
    WantDownSlow = 18,
    WantUpSlow = 19,
    Sprinting = 20,
    AscendBlock = 21,
    DescendBlock = 22,
    SneakToggleDown = 23,
    PersistSneak = 24,
    StartSprinting = 25,
    StopSprinting = 26,
    StartSneaking = 27,
    StopSneaking = 28,
    StartSwimming = 29,
    StopSwimming = 30,
    StartJumping = 31,
    StartGliding = 32,
    StopGliding = 33,
    PerformItemInteraction = 34,
    PerformBlockActions = 35,
    PerformItemStackRequest = 36,
    HandledTeleport = 37,
    Emoting = 38,
    MissedSwing = 39,
    StartCrawling = 40,
    StopCrawling = 41,
    StartFlying = 42,
    StopFlying = 43,
    ClientAckServerData = 44,
    IsInClientPredictedVehicle = 45,
    PaddlingLeft = 46,
    PaddlingRight = 47,
    BlockBreakingDelayEnabled = 48,
    HorizontalCollision = 49,
    VerticalCollision = 50,
    DownLeft = 51,
    DownRight = 52,
    StartUsingItem = 53,
    IsCameraRelativeMovementEnabled = 54,
    IsRotControlledByMoveDirection = 55,
    StartSpinAttack = 56,
    StopSpinAttack = 57,
    IsHotbarOnlyTouch = 58,
    JumpReleasedRaw = 59,
    JumpPressedRaw = 60,
    JumpCurrentRaw = 61,
    SneakReleasedRaw = 62,
    SneakPressedRaw = 63,
    SneakCurrentRaw = 64,
    InternalUpdate = 65,
}

impl TryFrom<i32> for PlayerAuthInputInputData {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Ascend),
            1 => Ok(Self::Descend),
            2 => Ok(Self::NorthJump),
            3 => Ok(Self::JumpDown),
            4 => Ok(Self::SprintDown),
            5 => Ok(Self::ChangeHeight),
            6 => Ok(Self::Jumping),
            7 => Ok(Self::AutoJumpingInWater),
            8 => Ok(Self::Sneaking),
            9 => Ok(Self::SneakDown),
            10 => Ok(Self::Up),
            11 => Ok(Self::Down),
            12 => Ok(Self::Left),
            13 => Ok(Self::Right),
            14 => Ok(Self::UpLeft),
            15 => Ok(Self::UpRight),
            16 => Ok(Self::WantUp),
            17 => Ok(Self::WantDown),
            18 => Ok(Self::WantDownSlow),
            19 => Ok(Self::WantUpSlow),
            20 => Ok(Self::Sprinting),
            21 => Ok(Self::AscendBlock),
            22 => Ok(Self::DescendBlock),
            23 => Ok(Self::SneakToggleDown),
            24 => Ok(Self::PersistSneak),
            25 => Ok(Self::StartSprinting),
            26 => Ok(Self::StopSprinting),
            27 => Ok(Self::StartSneaking),
            28 => Ok(Self::StopSneaking),
            29 => Ok(Self::StartSwimming),
            30 => Ok(Self::StopSwimming),
            31 => Ok(Self::StartJumping),
            32 => Ok(Self::StartGliding),
            33 => Ok(Self::StopGliding),
            34 => Ok(Self::PerformItemInteraction),
            35 => Ok(Self::PerformBlockActions),
            36 => Ok(Self::PerformItemStackRequest),
            37 => Ok(Self::HandledTeleport),
            38 => Ok(Self::Emoting),
            39 => Ok(Self::MissedSwing),
            40 => Ok(Self::StartCrawling),
            41 => Ok(Self::StopCrawling),
            42 => Ok(Self::StartFlying),
            43 => Ok(Self::StopFlying),
            44 => Ok(Self::ClientAckServerData),
            45 => Ok(Self::IsInClientPredictedVehicle),
            46 => Ok(Self::PaddlingLeft),
            47 => Ok(Self::PaddlingRight),
            48 => Ok(Self::BlockBreakingDelayEnabled),
            49 => Ok(Self::HorizontalCollision),
            50 => Ok(Self::VerticalCollision),
            51 => Ok(Self::DownLeft),
            52 => Ok(Self::DownRight),
            53 => Ok(Self::StartUsingItem),
            54 => Ok(Self::IsCameraRelativeMovementEnabled),
            55 => Ok(Self::IsRotControlledByMoveDirection),
            56 => Ok(Self::StartSpinAttack),
            57 => Ok(Self::StopSpinAttack),
            58 => Ok(Self::IsHotbarOnlyTouch),
            59 => Ok(Self::JumpReleasedRaw),
            60 => Ok(Self::JumpPressedRaw),
            61 => Ok(Self::JumpCurrentRaw),
            62 => Ok(Self::SneakReleasedRaw),
            63 => Ok(Self::SneakPressedRaw),
            64 => Ok(Self::SneakCurrentRaw),
            65 => Ok(Self::InternalUpdate),
            value => Err(value),
        }
    }
}

impl From<PlayerAuthInputInputData> for i32 {
    fn from(value: PlayerAuthInputInputData) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum PlayerLocationType {
    PlayerLocationCoordinates = 0,
}

impl TryFrom<i32> for PlayerLocationType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::PlayerLocationCoordinates),
            value => Err(value),
        }
    }
}

impl From<PlayerLocationType> for i32 {
    fn from(value: PlayerLocationType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i8)]
pub enum PlayerPermissionLevel {
    Visitor = 0,
    Member = 1,
    Operator = 2,
    Custom = 3,
}

impl TryFrom<i8> for PlayerPermissionLevel {
    type Error = i8;

    fn try_from(value: i8) -> Result<Self, i8> {
        match value {
            0 => Ok(Self::Visitor),
            1 => Ok(Self::Member),
            2 => Ok(Self::Operator),
            3 => Ok(Self::Custom),
            value => Err(value),
        }
    }
}

impl From<PlayerPermissionLevel> for i8 {
    fn from(value: PlayerPermissionLevel) -> Self {
        value as i8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum PlayerPositionModeComponentPositionMode {
    Normal = 0,
    Respawn = 1,
    Teleport = 2,
    OnlyHeadRot = 3,
}

impl TryFrom<u8> for PlayerPositionModeComponentPositionMode {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Normal),
            1 => Ok(Self::Respawn),
            2 => Ok(Self::Teleport),
            3 => Ok(Self::OnlyHeadRot),
            value => Err(value),
        }
    }
}

impl From<PlayerPositionModeComponentPositionMode> for u8 {
    fn from(value: PlayerPositionModeComponentPositionMode) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum PlayerRespawnState {
    SearchingForSpawn = 0,
    ReadyToSpawn = 1,
    ClientReadyToSpawn = 2,
}

impl TryFrom<u8> for PlayerRespawnState {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::SearchingForSpawn),
            1 => Ok(Self::ReadyToSpawn),
            2 => Ok(Self::ClientReadyToSpawn),
            value => Err(value),
        }
    }
}

impl From<PlayerRespawnState> for u8 {
    fn from(value: PlayerRespawnState) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum PositionTrackingDBClientRequestAction {
    Query = 0,
}

impl TryFrom<u8> for PositionTrackingDBClientRequestAction {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Query),
            value => Err(value),
        }
    }
}

impl From<PositionTrackingDBClientRequestAction> for u8 {
    fn from(value: PositionTrackingDBClientRequestAction) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum PositionTrackingDBServerBroadcastAction {
    Update = 0,
    Destroy = 1,
    NotFound = 2,
}

impl TryFrom<u8> for PositionTrackingDBServerBroadcastAction {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Update),
            1 => Ok(Self::Destroy),
            2 => Ok(Self::NotFound),
            value => Err(value),
        }
    }
}

impl From<PositionTrackingDBServerBroadcastAction> for u8 {
    fn from(value: PositionTrackingDBServerBroadcastAction) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum RandomDistributionType {
    SingleValued = 0,
    Uniform = 1,
    Gaussian = 2,
    InverseGaussian = 3,
    FixedGrid = 4,
    JitteredGrid = 5,
    Triangle = 6,
}

impl TryFrom<i32> for RandomDistributionType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::SingleValued),
            1 => Ok(Self::Uniform),
            2 => Ok(Self::Gaussian),
            3 => Ok(Self::InverseGaussian),
            4 => Ok(Self::FixedGrid),
            5 => Ok(Self::JitteredGrid),
            6 => Ok(Self::Triangle),
            value => Err(value),
        }
    }
}

impl From<RandomDistributionType> for i32 {
    fn from(value: RandomDistributionType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum RecipeUnlockingRequirementUnlockingContext {
    None = 0,
    AlwaysUnlocked = 1,
    PlayerInWater = 2,
    PlayerHasManyItems = 3,
}

impl TryFrom<i32> for RecipeUnlockingRequirementUnlockingContext {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::AlwaysUnlocked),
            2 => Ok(Self::PlayerInWater),
            3 => Ok(Self::PlayerHasManyItems),
            value => Err(value),
        }
    }
}

impl From<RecipeUnlockingRequirementUnlockingContext> for i32 {
    fn from(value: RecipeUnlockingRequirementUnlockingContext) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum RequestAbilityType {
    Unset = 0,
    Bool = 1,
    Float = 2,
}

impl TryFrom<u8> for RequestAbilityType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Unset),
            1 => Ok(Self::Bool),
            2 => Ok(Self::Float),
            value => Err(value),
        }
    }
}

impl From<RequestAbilityType> for u8 {
    fn from(value: RequestAbilityType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum RewindType {
    Player = 0,
    Vehicle = 1,
}

impl TryFrom<u8> for RewindType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Player),
            1 => Ok(Self::Vehicle),
            value => Err(value),
        }
    }
}

impl From<RewindType> for u8 {
    fn from(value: RewindType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum Rotation {
    None = 0,
    Rotate90 = 1,
    Rotate180 = 2,
    Rotate270 = 3,
}

impl TryFrom<u8> for Rotation {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::Rotate90),
            2 => Ok(Self::Rotate180),
            3 => Ok(Self::Rotate270),
            value => Err(value),
        }
    }
}

impl From<Rotation> for u8 {
    fn from(value: Rotation) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ScoreboardIdentityPacketType {
    Update = 0,
    Remove = 1,
}

impl TryFrom<u8> for ScoreboardIdentityPacketType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Update),
            1 => Ok(Self::Remove),
            value => Err(value),
        }
    }
}

impl From<ScoreboardIdentityPacketType> for u8 {
    fn from(value: ScoreboardIdentityPacketType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ScriptModuleMinecraftScriptPrimitiveShapeType {
    Line = 0,
    Box = 1,
    Sphere = 2,
    Circle = 3,
    Text = 4,
    Arrow = 5,
    Cylinder = 6,
    Pyramid = 7,
    Ellipsoid = 8,
    Cone = 9,
}

impl TryFrom<u8> for ScriptModuleMinecraftScriptPrimitiveShapeType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Line),
            1 => Ok(Self::Box),
            2 => Ok(Self::Sphere),
            3 => Ok(Self::Circle),
            4 => Ok(Self::Text),
            5 => Ok(Self::Arrow),
            6 => Ok(Self::Cylinder),
            7 => Ok(Self::Pyramid),
            8 => Ok(Self::Ellipsoid),
            9 => Ok(Self::Cone),
            value => Err(value),
        }
    }
}

impl From<ScriptModuleMinecraftScriptPrimitiveShapeType> for u8 {
    fn from(value: ScriptModuleMinecraftScriptPrimitiveShapeType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum ServerEditorConnectionPolicy {
    MatchWorldType = 0,
    EditorOnly = 1,
    VanillaOnly = 2,
    Mixed = 3,
}

impl TryFrom<i32> for ServerEditorConnectionPolicy {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::MatchWorldType),
            1 => Ok(Self::EditorOnly),
            2 => Ok(Self::VanillaOnly),
            3 => Ok(Self::Mixed),
            value => Err(value),
        }
    }
}

impl From<ServerEditorConnectionPolicy> for i32 {
    fn from(value: ServerEditorConnectionPolicy) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ServerWaypointGroupAction {
    None = 0,
    Add = 1,
    Remove = 2,
    Update = 3,
}

impl TryFrom<u8> for ServerWaypointGroupAction {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::Add),
            2 => Ok(Self::Remove),
            3 => Ok(Self::Update),
            value => Err(value),
        }
    }
}

impl From<ServerWaypointGroupAction> for u8 {
    fn from(value: ServerWaypointGroupAction) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum ServerboundLoadingScreenPacketType {
    StartLoadingScreen = 1,
    EndLoadingScreen = 2,
}

impl TryFrom<i32> for ServerboundLoadingScreenPacketType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            1 => Ok(Self::StartLoadingScreen),
            2 => Ok(Self::EndLoadingScreen),
            value => Err(value),
        }
    }
}

impl From<ServerboundLoadingScreenPacketType> for i32 {
    fn from(value: ServerboundLoadingScreenPacketType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum SetTitleTitleType {
    Clear = 0,
    Reset = 1,
    Title = 2,
    Subtitle = 3,
    Actionbar = 4,
    Times = 5,
    TitleTextObject = 6,
    SubtitleTextObject = 7,
    ActionbarTextObject = 8,
}

impl TryFrom<i32> for SetTitleTitleType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Clear),
            1 => Ok(Self::Reset),
            2 => Ok(Self::Title),
            3 => Ok(Self::Subtitle),
            4 => Ok(Self::Actionbar),
            5 => Ok(Self::Times),
            6 => Ok(Self::TitleTextObject),
            7 => Ok(Self::SubtitleTextObject),
            8 => Ok(Self::ActionbarTextObject),
            value => Err(value),
        }
    }
}

impl From<SetTitleTitleType> for i32 {
    fn from(value: SetTitleTitleType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum ShowStoreOfferRedirectType {
    MarketplaceOffer = 0,
    DressingRoomOffer = 1,
    ThirdPartyServerPage = 2,
}

impl TryFrom<u8> for ShowStoreOfferRedirectType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::MarketplaceOffer),
            1 => Ok(Self::DressingRoomOffer),
            2 => Ok(Self::ThirdPartyServerPage),
            value => Err(value),
        }
    }
}

impl From<ShowStoreOfferRedirectType> for u8 {
    fn from(value: ShowStoreOfferRedirectType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u16)]
pub enum SimpleEventSubtype {
    UninitializedSubtype = 0,
    EnableCommands = 1,
    DisableCommands = 2,
    UnlockWorldTemplateSettings = 3,
}

impl TryFrom<u16> for SimpleEventSubtype {
    type Error = u16;

    fn try_from(value: u16) -> Result<Self, u16> {
        match value {
            0 => Ok(Self::UninitializedSubtype),
            1 => Ok(Self::EnableCommands),
            2 => Ok(Self::DisableCommands),
            3 => Ok(Self::UnlockWorldTemplateSettings),
            value => Err(value),
        }
    }
}

impl From<SimpleEventSubtype> for u16 {
    fn from(value: SimpleEventSubtype) -> Self {
        value as u16
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum SimulationTypeType {
    Game = 0,
    Editor = 1,
    Test = 2,
    Invalid = 3,
}

impl TryFrom<u8> for SimulationTypeType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Game),
            1 => Ok(Self::Editor),
            2 => Ok(Self::Test),
            3 => Ok(Self::Invalid),
            value => Err(value),
        }
    }
}

impl From<SimulationTypeType> for u8 {
    fn from(value: SimulationTypeType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum SocialGamePublishSetting {
    NoMultiPlay = 0,
    InviteOnly = 1,
    FriendsOnly = 2,
    FriendsOfFriends = 3,
    Public = 4,
}

impl TryFrom<i32> for SocialGamePublishSetting {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::NoMultiPlay),
            1 => Ok(Self::InviteOnly),
            2 => Ok(Self::FriendsOnly),
            3 => Ok(Self::FriendsOfFriends),
            4 => Ok(Self::Public),
            value => Err(value),
        }
    }
}

impl From<SocialGamePublishSetting> for i32 {
    fn from(value: SocialGamePublishSetting) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum SoftEnumUpdateType {
    Add = 0,
    Remove = 1,
    Replace = 2,
}

impl TryFrom<u8> for SoftEnumUpdateType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Add),
            1 => Ok(Self::Remove),
            2 => Ok(Self::Replace),
            value => Err(value),
        }
    }
}

impl From<SoftEnumUpdateType> for u8 {
    fn from(value: SoftEnumUpdateType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i16)]
pub enum SpawnBiomeType {
    Default = 0,
    UserDefined = 1,
}

impl TryFrom<i16> for SpawnBiomeType {
    type Error = i16;

    fn try_from(value: i16) -> Result<Self, i16> {
        match value {
            0 => Ok(Self::Default),
            1 => Ok(Self::UserDefined),
            value => Err(value),
        }
    }
}

impl From<SpawnBiomeType> for i16 {
    fn from(value: SpawnBiomeType) -> Self {
        value as i16
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum SpawnPositionType {
    PlayerRespawn = 0,
    WorldSpawn = 1,
}

impl TryFrom<i32> for SpawnPositionType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::PlayerRespawn),
            1 => Ok(Self::WorldSpawn),
            value => Err(value),
        }
    }
}

impl From<SpawnPositionType> for i32 {
    fn from(value: SpawnPositionType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum StructureBlockType {
    Data = 0,
    Save = 1,
    Load = 2,
    Corner = 3,
    Invalid = 4,
    Export = 5,
}

impl TryFrom<i32> for StructureBlockType {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            0 => Ok(Self::Data),
            1 => Ok(Self::Save),
            2 => Ok(Self::Load),
            3 => Ok(Self::Corner),
            4 => Ok(Self::Invalid),
            5 => Ok(Self::Export),
            value => Err(value),
        }
    }
}

impl From<StructureBlockType> for i32 {
    fn from(value: StructureBlockType) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum StructureRedstoneSaveMode {
    SavesToMemory = 0,
    SavesToDisk = 1,
}

impl TryFrom<u8> for StructureRedstoneSaveMode {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::SavesToMemory),
            1 => Ok(Self::SavesToDisk),
            value => Err(value),
        }
    }
}

impl From<StructureRedstoneSaveMode> for u8 {
    fn from(value: StructureRedstoneSaveMode) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum StructureTemplateRequestOperation {
    None = 0,
    ExportFromSaveMode = 1,
    ExportFromLoadMode = 2,
    QuerySavedStructure = 3,
}

impl TryFrom<u8> for StructureTemplateRequestOperation {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::ExportFromSaveMode),
            2 => Ok(Self::ExportFromLoadMode),
            3 => Ok(Self::QuerySavedStructure),
            value => Err(value),
        }
    }
}

impl From<StructureTemplateRequestOperation> for u8 {
    fn from(value: StructureTemplateRequestOperation) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum StructureTemplateResponseType {
    None = 0,
    Export = 1,
    Query = 2,
}

impl TryFrom<u8> for StructureTemplateResponseType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::None),
            1 => Ok(Self::Export),
            2 => Ok(Self::Query),
            value => Err(value),
        }
    }
}

impl From<StructureTemplateResponseType> for u8 {
    fn from(value: StructureTemplateResponseType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum SubChunkHeightMapDataType {
    NoData = 0,
    HasData = 1,
    AllTooHigh = 2,
    AllTooLow = 3,
}

impl TryFrom<u8> for SubChunkHeightMapDataType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::NoData),
            1 => Ok(Self::HasData),
            2 => Ok(Self::AllTooHigh),
            3 => Ok(Self::AllTooLow),
            value => Err(value),
        }
    }
}

impl From<SubChunkHeightMapDataType> for u8 {
    fn from(value: SubChunkHeightMapDataType) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum SubChunkSubChunkRequestResult {
    Success = 1,
    LevelChunkDoesntExist = 2,
    WrongDimension = 3,
    PlayerDoesntExist = 4,
    IndexOutOfBounds = 5,
    SuccessAllAir = 6,
}

impl TryFrom<u8> for SubChunkSubChunkRequestResult {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            1 => Ok(Self::Success),
            2 => Ok(Self::LevelChunkDoesntExist),
            3 => Ok(Self::WrongDimension),
            4 => Ok(Self::PlayerDoesntExist),
            5 => Ok(Self::IndexOutOfBounds),
            6 => Ok(Self::SuccessAllAir),
            value => Err(value),
        }
    }
}

impl From<SubChunkSubChunkRequestResult> for u8 {
    fn from(value: SubChunkSubChunkRequestResult) -> Self {
        value as u8
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(i32)]
pub enum TextProcessingEventOrigin {
    Unknown = -1,
    ServerChatPublic = 0,
    ServerChatWhisper = 1,
    SignText = 2,
    AnvilText = 3,
    BookAndQuillText = 4,
    CommandBlockText = 5,
    BlockActorDataText = 6,
    JoinEventText = 7,
    LeaveEventText = 8,
    SlashCommandChat = 9,
    CartographyText = 10,
    KickCommand = 11,
    TitleCommand = 12,
    SummonCommand = 13,
    ServerForm = 14,
    DataDrivenUi = 15,
}

impl TryFrom<i32> for TextProcessingEventOrigin {
    type Error = i32;

    fn try_from(value: i32) -> Result<Self, i32> {
        match value {
            -1 => Ok(Self::Unknown),
            0 => Ok(Self::ServerChatPublic),
            1 => Ok(Self::ServerChatWhisper),
            2 => Ok(Self::SignText),
            3 => Ok(Self::AnvilText),
            4 => Ok(Self::BookAndQuillText),
            5 => Ok(Self::CommandBlockText),
            6 => Ok(Self::BlockActorDataText),
            7 => Ok(Self::JoinEventText),
            8 => Ok(Self::LeaveEventText),
            9 => Ok(Self::SlashCommandChat),
            10 => Ok(Self::CartographyText),
            11 => Ok(Self::KickCommand),
            12 => Ok(Self::TitleCommand),
            13 => Ok(Self::SummonCommand),
            14 => Ok(Self::ServerForm),
            15 => Ok(Self::DataDrivenUi),
            value => Err(value),
        }
    }
}

impl From<TextProcessingEventOrigin> for i32 {
    fn from(value: TextProcessingEventOrigin) -> Self {
        value as i32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u32)]
pub enum UnlockedRecipesPacketType {
    Empty = 0,
    Initially = 1,
    Newly = 2,
    Remove = 3,
    RemoveAll = 4,
}

impl TryFrom<u32> for UnlockedRecipesPacketType {
    type Error = u32;

    fn try_from(value: u32) -> Result<Self, u32> {
        match value {
            0 => Ok(Self::Empty),
            1 => Ok(Self::Initially),
            2 => Ok(Self::Newly),
            3 => Ok(Self::Remove),
            4 => Ok(Self::RemoveAll),
            value => Err(value),
        }
    }
}

impl From<UnlockedRecipesPacketType> for u32 {
    fn from(value: UnlockedRecipesPacketType) -> Self {
        value as u32
    }
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
#[repr(u8)]
pub enum VillageType {
    Desert = 0,
    Ice = 1,
    Savanna = 2,
    Taiga = 3,
    Default = 4,
}

impl TryFrom<u8> for VillageType {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, u8> {
        match value {
            0 => Ok(Self::Desert),
            1 => Ok(Self::Ice),
            2 => Ok(Self::Savanna),
            3 => Ok(Self::Taiga),
            4 => Ok(Self::Default),
            value => Err(value),
        }
    }
}

impl From<VillageType> for u8 {
    fn from(value: VillageType) -> Self {
        value as u8
    }
}
