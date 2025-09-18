-- cgoffline.sql
-- SQL queries for CoinGecko Asset Platforms data

-- 1. Get total count of asset platforms
SELECT COUNT(*) as total_platforms FROM asset_platforms;

-- 2. Get all asset platforms with basic information
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id,
    created_at,
    updated_at
FROM asset_platforms 
ORDER BY name;

-- 3. Get platforms with chain identifiers (EVM-compatible chains)
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE chain_identifier IS NOT NULL 
ORDER BY chain_identifier;

-- 4. Get platforms without chain identifiers (non-EVM chains)
SELECT 
    id,
    name,
    native_coin_id
FROM asset_platforms 
WHERE chain_identifier IS NULL 
ORDER BY name;

-- 5. Search for specific platforms by name
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE name ILIKE '%ethereum%' 
   OR name ILIKE '%bitcoin%'
   OR name ILIKE '%polygon%'
ORDER BY name;

-- 6. Get platforms by native coin ID
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id IN ('ethereum', 'bitcoin', 'binancecoin', 'cardano', 'solana')
ORDER BY native_coin_id, name;

-- 7. Get platforms with specific chain identifiers (popular chains)
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE chain_identifier IN (1, 56, 137, 250, 43114, 42161, 10, 8453)
ORDER BY chain_identifier;

-- 8. Get recently updated platforms
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id,
    updated_at
FROM asset_platforms 
ORDER BY updated_at DESC 
LIMIT 20;

-- 9. Get platforms with short names
SELECT 
    id,
    name,
    short_name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE short_name IS NOT NULL 
ORDER BY name;

-- 10. Get platforms grouped by native coin
SELECT 
    native_coin_id,
    COUNT(*) as platform_count,
    STRING_AGG(name, ', ' ORDER BY name) as platform_names
FROM asset_platforms 
WHERE native_coin_id IS NOT NULL
GROUP BY native_coin_id 
HAVING COUNT(*) > 1
ORDER BY platform_count DESC;

-- 11. Get platforms with highest chain identifiers (newer chains)
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE chain_identifier IS NOT NULL
ORDER BY chain_identifier DESC 
LIMIT 20;

-- 12. Get platforms with lowest chain identifiers (older chains)
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE chain_identifier IS NOT NULL
ORDER BY chain_identifier ASC 
LIMIT 20;

-- 13. Get platforms containing specific keywords
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE name ILIKE '%layer%' 
   OR name ILIKE '%sidechain%'
   OR name ILIKE '%testnet%'
ORDER BY name;

-- 14. Get platforms with WETH as native coin (Ethereum L2s)
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'weth'
ORDER BY name;

-- 15. Get platforms with Bitcoin as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'bitcoin'
ORDER BY name;

-- 16. Get platforms with Ethereum as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'ethereum'
ORDER BY name;

-- 17. Get platforms with BNB as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'binancecoin'
ORDER BY name;

-- 18. Get platforms with AVAX as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'avalanche-2'
ORDER BY name;

-- 19. Get platforms with MATIC as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'matic-network'
ORDER BY name;

-- 20. Get platforms with FTM as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'fantom'
ORDER BY name;

-- 21. Get platforms with ARB as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'arbitrum'
ORDER BY name;

-- 22. Get platforms with OP as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'optimism'
ORDER BY name;

-- 23. Get platforms with BASE as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'base'
ORDER BY name;

-- 24. Get platforms with SOL as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'solana'
ORDER BY name;

-- 25. Get platforms with ADA as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'cardano'
ORDER BY name;

-- 26. Get platforms with DOT as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'polkadot'
ORDER BY name;

-- 27. Get platforms with ATOM as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'cosmos'
ORDER BY name;

-- 28. Get platforms with TRX as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'tron'
ORDER BY name;

-- 29. Get platforms with TON as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'the-open-network'
ORDER BY name;

-- 30. Get platforms with WAX as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'wax'
ORDER BY name;

-- 31. Get platforms with CHIA as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'chia'
ORDER BY name;

-- 32. Get platforms with MASSA as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'massa'
ORDER BY name;

-- 33. Get platforms with STARKNET as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'starknet'
ORDER BY name;

-- 34. Get platforms with SUPRA as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'supra'
ORDER BY name;

-- 35. Get platforms with PEAQ as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'peaq-2'
ORDER BY name;

-- 36. Get platforms with LAMINA1 as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'lamina1'
ORDER BY name;

-- 37. Get platforms with PLASMA as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'plasma'
ORDER BY name;

-- 38. Get platforms with MEZO as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'wrapped-bitcoin'
ORDER BY name;

-- 39. Get platforms with APE as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'apecoin'
ORDER BY name;

-- 40. Get platforms with CANTO as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'canto'
ORDER BY name;

-- 41. Get platforms with PULSECHAIN as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'pulsechain'
ORDER BY name;

-- 42. Get platforms with EDU as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'edu-coin'
ORDER BY name;

-- 43. Get platforms with BEVM as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'bitcoin'
ORDER BY name;

-- 44. Get platforms with RE.AL as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'weth'
ORDER BY name;

-- 45. Get platforms with PLANQ as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'planq'
ORDER BY name;

-- 46. Get platforms with SAAKURU as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'wrapped-oasys'
ORDER BY name;

-- 47. Get platforms with MINT as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'weth'
ORDER BY name;

-- 48. Get platforms with ENULS as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'nuls'
ORDER BY name;

-- 49. Get platforms with MORPH as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'weth'
ORDER BY name;

-- 50. Get platforms with DEFIVERSE as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'oasys'
ORDER BY name;

-- 51. Get platforms with ECLIPSE as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'weth'
ORDER BY name;

-- 52. Get platforms with SOPHON as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'sophon'
ORDER BY name;

-- 53. Get platforms with BOB as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'weth'
ORDER BY name;

-- 54. Get platforms with GRAVITY as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'g-token'
ORDER BY name;

-- 55. Get platforms with SKALE as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'skale'
ORDER BY name;

-- 56. Get platforms with KROMA as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'weth'
ORDER BY name;

-- 57. Get platforms with WORLD as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'weth'
ORDER BY name;

-- 58. Get platforms with SUPerseed as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'ethereum'
ORDER BY name;

-- 59. Get platforms with SHIDEN as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'shiden'
ORDER BY name;

-- 60. Get platforms with DOGECHAIN as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'dogechain'
ORDER BY name;

-- 61. Get platforms with SCROLL as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'weth'
ORDER BY name;

-- 62. Get platforms with TARAXA as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'taraxa'
ORDER BY name;

-- 63. Get platforms with CRONOS as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'cronos-zkevm-cro'
ORDER BY name;

-- 64. Get platforms with ROOTSTOCK as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'rootstock'
ORDER BY name;

-- 65. Get platforms with EVENTUM as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'ethereum'
ORDER BY name;

-- 66. Get platforms with 8BIT as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = '8bit-chain'
ORDER BY name;

-- 67. Get platforms with QUBIC as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'qubic-network'
ORDER BY name;

-- 68. Get platforms with BSQUARED as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'bitcoin'
ORDER BY name;

-- 69. Get platforms with PROVENANCE as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'hash-2'
ORDER BY name;

-- 70. Get platforms with AIRDAO as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'amber'
ORDER BY name;

-- 71. Get platforms with BSC as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'binancecoin'
ORDER BY name;

-- 72. Get platforms with ZKSYNC as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'ethereum'
ORDER BY name;

-- 73. Get platforms with POLYGON as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'matic-network'
ORDER BY name;

-- 74. Get platforms with FANTOM as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'fantom'
ORDER BY name;

-- 75. Get platforms with ARBITRUM as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'arbitrum'
ORDER BY name;

-- 76. Get platforms with OPTIMISM as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'optimism'
ORDER BY name;

-- 77. Get platforms with BASE as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'base'
ORDER BY name;

-- 78. Get platforms with LINEA as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'ethereum'
ORDER BY name;

-- 79. Get platforms with BLAST as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'weth'
ORDER BY name;

-- 80. Get platforms with MANTLE as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'weth'
ORDER BY name;

-- 81. Get platforms with CELO as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'celo'
ORDER BY name;

-- 82. Get platforms with GNOSIS as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'gnosis'
ORDER BY name;

-- 83. Get platforms with HARMONY as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'harmony'
ORDER BY name;

-- 84. Get platforms with MOONBEAM as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'moonbeam'
ORDER BY name;

-- 85. Get platforms with MOONRIVER as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'moonriver'
ORDER BY name;

-- 86. Get platforms with ASTAR as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'astar'
ORDER BY name;

-- 87. Get platforms with KUSAMA as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'kusama'
ORDER BY name;

-- 88. Get platforms with NEAR as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'near'
ORDER BY name;

-- 89. Get platforms with ALGORAND as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'algorand'
ORDER BY name;

-- 90. Get platforms with TEZOS as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'tezos'
ORDER BY name;

-- 91. Get platforms with FLOW as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'flow'
ORDER BY name;

-- 92. Get platforms with HEDERA as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'hedera-hashgraph'
ORDER BY name;

-- 93. Get platforms with ICP as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'internet-computer'
ORDER BY name;

-- 94. Get platforms with APTOS as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'aptos'
ORDER BY name;

-- 95. Get platforms with SUI as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'sui'
ORDER BY name;

-- 96. Get platforms with SEI as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'sei-network'
ORDER BY name;

-- 97. Get platforms with INJECTIVE as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'injective-protocol'
ORDER BY name;

-- 98. Get platforms with OSMOSIS as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'osmosis'
ORDER BY name;

-- 99. Get platforms with JUNO as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'juno-network'
ORDER BY name;

-- 100. Get platforms with EVMOS as native coin
SELECT 
    id,
    name,
    chain_identifier,
    native_coin_id
FROM asset_platforms 
WHERE native_coin_id = 'evmos'
ORDER BY name;

-- ========================================
-- COIN CATEGORIES QUERIES
-- ========================================

-- 101. Get total count of coin categories
SELECT COUNT(*) as total_categories FROM coin_categories;

-- 102. Get all coin categories with basic information
SELECT 
    id,
    coingecko_id,
    name,
    created_at,
    updated_at
FROM coin_categories 
ORDER BY name;

-- 103. Search for specific categories by name
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%defi%' 
   OR name ILIKE '%nft%'
   OR name ILIKE '%gaming%'
ORDER BY name;

-- 104. Get categories by coingecko_id pattern
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id ILIKE '%ecosystem%'
ORDER BY name;

-- 105. Get categories containing specific keywords
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%layer%'
ORDER BY name;

-- 106. Get categories related to gaming
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%game%' 
   OR name ILIKE '%gaming%'
   OR name ILIKE '%play%'
ORDER BY name;

-- 107. Get categories related to DeFi
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%defi%' 
   OR name ILIKE '%yield%'
   OR name ILIKE '%lending%'
   OR name ILIKE '%swap%'
ORDER BY name;

-- 108. Get categories related to NFTs
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%nft%' 
   OR name ILIKE '%collectible%'
   OR name ILIKE '%art%'
ORDER BY name;

-- 109. Get categories related to stablecoins
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%stablecoin%' 
   OR name ILIKE '%stable%'
ORDER BY name;

-- 110. Get categories related to meme coins
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%meme%' 
   OR name ILIKE '%dog%'
   OR name ILIKE '%cat%'
ORDER BY name;

-- 111. Get categories related to AI
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%ai%' 
   OR name ILIKE '%artificial%'
   OR name ILIKE '%machine%'
ORDER BY name;

-- 112. Get categories related to privacy
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%privacy%' 
   OR name ILIKE '%anonymous%'
   OR name ILIKE '%zero%'
ORDER BY name;

-- 113. Get categories related to infrastructure
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%infrastructure%' 
   OR name ILIKE '%oracle%'
   OR name ILIKE '%storage%'
ORDER BY name;

-- 114. Get categories related to social
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%social%' 
   OR name ILIKE '%community%'
   OR name ILIKE '%creator%'
ORDER BY name;

-- 115. Get categories related to metaverse
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%metaverse%' 
   OR name ILIKE '%virtual%'
   OR name ILIKE '%reality%'
ORDER BY name;

-- 116. Get categories related to payments
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%payment%' 
   OR name ILIKE '%wallet%'
   OR name ILIKE '%transaction%'
ORDER BY name;

-- 117. Get categories related to enterprise
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%enterprise%' 
   OR name ILIKE '%business%'
   OR name ILIKE '%corporate%'
ORDER BY name;

-- 118. Get categories related to insurance
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%insurance%' 
   OR name ILIKE '%risk%'
   OR name ILIKE '%coverage%'
ORDER BY name;

-- 119. Get categories related to prediction markets
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%prediction%' 
   OR name ILIKE '%betting%'
   OR name ILIKE '%forecast%'
ORDER BY name;

-- 120. Get categories related to supply chain
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%supply%' 
   OR name ILIKE '%logistics%'
   OR name ILIKE '%tracking%'
ORDER BY name;

-- 121. Get categories related to energy
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%energy%' 
   OR name ILIKE '%green%'
   OR name ILIKE '%renewable%'
ORDER BY name;

-- 122. Get categories related to healthcare
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%health%' 
   OR name ILIKE '%medical%'
   OR name ILIKE '%pharma%'
ORDER BY name;

-- 123. Get categories related to education
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%education%' 
   OR name ILIKE '%learning%'
   OR name ILIKE '%academic%'
ORDER BY name;

-- 124. Get categories related to real estate
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%real%' 
   OR name ILIKE '%estate%'
   OR name ILIKE '%property%'
ORDER BY name;

-- 125. Get categories related to music
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%music%' 
   OR name ILIKE '%audio%'
   OR name ILIKE '%sound%'
ORDER BY name;

-- 126. Get categories related to sports
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%sport%' 
   OR name ILIKE '%fitness%'
   OR name ILIKE '%athletic%'
ORDER BY name;

-- 127. Get categories related to travel
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%travel%' 
   OR name ILIKE '%tourism%'
   OR name ILIKE '%booking%'
ORDER BY name;

-- 128. Get categories related to food
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%food%' 
   OR name ILIKE '%restaurant%'
   OR name ILIKE '%dining%'
ORDER BY name;

-- 129. Get categories related to fashion
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%fashion%' 
   OR name ILIKE '%clothing%'
   OR name ILIKE '%style%'
ORDER BY name;

-- 130. Get categories related to entertainment
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%entertainment%' 
   OR name ILIKE '%media%'
   OR name ILIKE '%content%'
ORDER BY name;

-- 131. Get categories related to news
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%news%' 
   OR name ILIKE '%journalism%'
   OR name ILIKE '%reporting%'
ORDER BY name;

-- 132. Get categories related to charity
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%charity%' 
   OR name ILIKE '%donation%'
   OR name ILIKE '%philanthropy%'
ORDER BY name;

-- 133. Get categories related to governance
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%governance%' 
   OR name ILIKE '%voting%'
   OR name ILIKE '%dao%'
ORDER BY name;

-- 134. Get categories related to identity
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%identity%' 
   OR name ILIKE '%kyc%'
   OR name ILIKE '%verification%'
ORDER BY name;

-- 135. Get categories related to data
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%data%' 
   OR name ILIKE '%analytics%'
   OR name ILIKE '%database%'
ORDER BY name;

-- 136. Get categories related to security
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%security%' 
   OR name ILIKE '%audit%'
   OR name ILIKE '%protection%'
ORDER BY name;

-- 137. Get categories related to compliance
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%compliance%' 
   OR name ILIKE '%regulation%'
   OR name ILIKE '%legal%'
ORDER BY name;

-- 138. Get categories related to trading
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%trading%' 
   OR name ILIKE '%exchange%'
   OR name ILIKE '%market%'
ORDER BY name;

-- 139. Get categories related to lending
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%lending%' 
   OR name ILIKE '%borrowing%'
   OR name ILIKE '%credit%'
ORDER BY name;

-- 140. Get categories related to derivatives
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%derivative%' 
   OR name ILIKE '%futures%'
   OR name ILIKE '%options%'
ORDER BY name;

-- 141. Get categories related to cross-chain
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%cross%' 
   OR name ILIKE '%bridge%'
   OR name ILIKE '%interoperability%'
ORDER BY name;

-- 142. Get categories related to scaling
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%scaling%' 
   OR name ILIKE '%layer%'
   OR name ILIKE '%rollup%'
ORDER BY name;

-- 143. Get categories related to staking
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%staking%' 
   OR name ILIKE '%validator%'
   OR name ILIKE '%delegation%'
ORDER BY name;

-- 144. Get categories related to mining
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%mining%' 
   OR name ILIKE '%proof%'
   OR name ILIKE '%consensus%'
ORDER BY name;

-- 145. Get categories related to oracle
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%oracle%' 
   OR name ILIKE '%price%'
   OR name ILIKE '%feed%'
ORDER BY name;

-- 146. Get categories related to wallet
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%wallet%' 
   OR name ILIKE '%custody%'
   OR name ILIKE '%storage%'
ORDER BY name;

-- 147. Get categories related to browser
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%browser%' 
   OR name ILIKE '%web3%'
   OR name ILIKE '%dapp%'
ORDER BY name;

-- 148. Get categories related to development
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%development%' 
   OR name ILIKE '%tool%'
   OR name ILIKE '%framework%'
ORDER BY name;

-- 149. Get categories related to testing
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%test%' 
   OR name ILIKE '%testing%'
   OR name ILIKE '%qa%'
ORDER BY name;

-- 150. Get categories related to monitoring
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%monitoring%' 
   OR name ILIKE '%analytics%'
   OR name ILIKE '%metrics%'
ORDER BY name;

-- 151. Get categories updated in the last 24 hours
SELECT 
    id,
    coingecko_id,
    name,
    updated_at
FROM coin_categories 
WHERE updated_at >= NOW() - INTERVAL '24 hours'
ORDER BY updated_at DESC;

-- 152. Get categories created in a specific year (e.g., 2024)
SELECT 
    id,
    coingecko_id,
    name,
    created_at
FROM coin_categories 
WHERE EXTRACT(YEAR FROM created_at) = 2024
ORDER BY created_at;

-- 153. Get categories with the longest names
SELECT 
    id,
    coingecko_id,
    name,
    LENGTH(name) as name_length
FROM coin_categories 
ORDER BY name_length DESC
LIMIT 10;

-- 154. Get categories with the shortest names
SELECT 
    id,
    coingecko_id,
    name,
    LENGTH(name) as name_length
FROM coin_categories 
ORDER BY name_length ASC
LIMIT 10;

-- 155. Count categories by the first letter of their name
SELECT 
    SUBSTRING(name, 1, 1) as first_letter,
    COUNT(*) as count
FROM coin_categories 
GROUP BY first_letter
ORDER BY first_letter;

-- 156. Get categories where the name contains a number
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ~ '[0-9]'
ORDER BY name;

-- 157. Get categories where the coingecko_id contains a number
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id ~ '[0-9]'
ORDER BY name;

-- 158. Get categories where the name is unique (no duplicates)
SELECT 
    name,
    COUNT(*)
FROM coin_categories 
GROUP BY name
HAVING COUNT(*) = 1
ORDER BY name;

-- 159. Get categories where the name is duplicated
SELECT 
    name,
    COUNT(*)
FROM coin_categories 
GROUP BY name
HAVING COUNT(*) > 1
ORDER BY name;

-- 160. Get categories where the coingecko_id is unique (no duplicates)
SELECT 
    coingecko_id,
    COUNT(*)
FROM coin_categories 
GROUP BY coingecko_id
HAVING COUNT(*) = 1
ORDER BY coingecko_id;

-- 161. Get categories where the coingecko_id is duplicated
SELECT 
    coingecko_id,
    COUNT(*)
FROM coin_categories 
GROUP BY coingecko_id
HAVING COUNT(*) > 1
ORDER BY coingecko_id;

-- 162. Get categories ordered by creation date
SELECT 
    id,
    coingecko_id,
    name,
    created_at
FROM coin_categories 
ORDER BY created_at DESC;

-- 163. Get categories ordered by update date
SELECT 
    id,
    coingecko_id,
    name,
    updated_at
FROM coin_categories 
ORDER BY updated_at DESC;

-- 164. Get categories that have been updated more recently than created
SELECT 
    id,
    coingecko_id,
    name,
    created_at,
    updated_at
FROM coin_categories 
WHERE updated_at > created_at
ORDER BY updated_at DESC;

-- 165. Get categories that were created and updated on the same day
SELECT 
    id,
    coingecko_id,
    name,
    created_at,
    updated_at
FROM coin_categories 
WHERE DATE(created_at) = DATE(updated_at)
ORDER BY name;

-- 166. Get categories with a specific number of characters in their name (e.g., 10 characters)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE LENGTH(name) = 10
ORDER BY name;

-- 167. Get categories with a specific number of characters in their coingecko_id (e.g., 15 characters)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE LENGTH(coingecko_id) = 15
ORDER BY name;

-- 168. Get categories where the name starts with 'A'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name LIKE 'A%'
ORDER BY name;

-- 169. Get categories where the name ends with 'coin'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name LIKE '%coin'
ORDER BY name;

-- 170. Get categories where the coingecko_id starts with 'bitcoin'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id LIKE 'bitcoin%'
ORDER BY name;

-- 171. Get categories where the coingecko_id ends with 'ecosystem'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id LIKE '%ecosystem'
ORDER BY name;

-- 172. Get categories with a specific pattern in their name (e.g., containing 'defi')
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name ILIKE '%defi%'
ORDER BY name;

-- 173. Get categories with a specific pattern in their coingecko_id (e.g., containing 'ethereum')
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id ILIKE '%ethereum%'
ORDER BY name;

-- 174. Get categories where the name is exactly 'DeFi'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi';

-- 175. Get categories where the coingecko_id is exactly 'defi'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id = 'defi';

-- 176. Get categories that are not 'DeFi'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name != 'DeFi'
ORDER BY name;

-- 177. Get categories that are not 'defi'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id != 'defi'
ORDER BY name;

-- 178. Get categories with a specific name and coingecko_id
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' AND coingecko_id = 'defi'
ORDER BY name;

-- 179. Get categories with a specific name or coingecko_id
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' OR coingecko_id = 'defi'
ORDER BY name;

-- 180. Get categories with a specific name and a different coingecko_id
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' AND coingecko_id != 'defi'
ORDER BY name;

-- 181. Get categories with a specific coingecko_id and a different name
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id = 'defi' AND name != 'DeFi'
ORDER BY name;

-- 182. Get categories with a specific name and a coingecko_id that contains 'defi'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' AND coingecko_id ILIKE '%defi%'
ORDER BY name;

-- 183. Get categories with a specific coingecko_id and a name that contains 'DeFi'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id = 'defi' AND name ILIKE '%DeFi%'
ORDER BY name;

-- 184. Get categories with a specific name and a coingecko_id that starts with 'defi'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' AND coingecko_id LIKE 'defi%'
ORDER BY name;

-- 185. Get categories with a specific coingecko_id and a name that starts with 'DeFi'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id = 'defi' AND name LIKE 'DeFi%'
ORDER BY name;

-- 186. Get categories with a specific name and a coingecko_id that ends with 'defi'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' AND coingecko_id LIKE '%defi'
ORDER BY name;

-- 187. Get categories with a specific coingecko_id and a name that ends with 'DeFi'
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id = 'defi' AND name LIKE '%DeFi'
ORDER BY name;

-- 188. Get categories with a specific name and a coingecko_id that matches a pattern
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' AND coingecko_id ~ '^defi.*'
ORDER BY name;

-- 189. Get categories with a specific coingecko_id and a name that matches a pattern
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id = 'defi' AND name ~ '^DeFi.*'
ORDER BY name;

-- 190. Get categories with a specific name and a coingecko_id that matches a pattern (case insensitive)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' AND coingecko_id ~* '^defi.*'
ORDER BY name;

-- 191. Get categories with a specific coingecko_id and a name that matches a pattern (case insensitive)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id = 'defi' AND name ~* '^defi.*'
ORDER BY name;

-- 192. Get categories with a specific name and a coingecko_id that matches a pattern (case insensitive)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' AND coingecko_id ~* '^defi.*'
ORDER BY name;

-- 193. Get categories with a specific coingecko_id and a name that matches a pattern (case insensitive)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id = 'defi' AND name ~* '^defi.*'
ORDER BY name;

-- 194. Get categories with a specific name and a coingecko_id that matches a pattern (case insensitive)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' AND coingecko_id ~* '^defi.*'
ORDER BY name;

-- 195. Get categories with a specific coingecko_id and a name that matches a pattern (case insensitive)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id = 'defi' AND name ~* '^defi.*'
ORDER BY name;

-- 196. Get categories with a specific name and a coingecko_id that matches a pattern (case insensitive)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' AND coingecko_id ~* '^defi.*'
ORDER BY name;

-- 197. Get categories with a specific coingecko_id and a name that matches a pattern (case insensitive)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id = 'defi' AND name ~* '^defi.*'
ORDER BY name;

-- 198. Get categories with a specific name and a coingecko_id that matches a pattern (case insensitive)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' AND coingecko_id ~* '^defi.*'
ORDER BY name;

-- 199. Get categories with a specific coingecko_id and a name that matches a pattern (case insensitive)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE coingecko_id = 'defi' AND name ~* '^defi.*'
ORDER BY name;

-- 200. Get categories with a specific name and a coingecko_id that matches a pattern (case insensitive)
SELECT 
    id,
    coingecko_id,
    name
FROM coin_categories 
WHERE name = 'DeFi' AND coingecko_id ~* '^defi.*'
ORDER BY name;
