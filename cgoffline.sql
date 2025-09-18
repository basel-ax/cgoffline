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
