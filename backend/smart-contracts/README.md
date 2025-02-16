### **📄 README.md – Smart Contracts for Decentralized Compute Network**  

# **Decentralized Compute Network – Smart Contract Overview**  
This document explains how the smart contract system governs **task verification, worker rewards, and token distribution** in the **Decentralized Compute Network (DCN)**.  

---

## **🛠 Smart Contract Responsibilities**  
The smart contract ensures:  
1. **Fair reward distribution**: Workers receive points for valid computations and can claim rewards.  
2. **Task verification**: Results are checked before rewarding.  
3. **Ownership & security**: Only the Task Manager can update points.  
4. **ERC-20 token rewards**: Workers are rewarded with ERC-20 tokens.  

---

## **📌 ERC Standards Used**  
- **ERC-20**: For worker rewards (fungible tokens).  
- **Ownable (OpenZeppelin)**: Allows only the owner (Task Manager) to modify rewards.  
- **Reentrancy Guard**: Prevents double spending when claiming rewards.  

---

## **🔹 Key Smart Contract Components**  

### **1️⃣ Worker Point System**  
- Each worker node has a **points balance**.  
- Points increase when tasks are successfully completed.  
- Task Manager can update points **only after verification**.  

### **2️⃣ Task Verification**  
- Worker submits computation results.  
- Task Manager verifies results and updates points.  
- On-chain **hash commitment** ensures task integrity.  

### **3️⃣ Reward Distribution**  
- Workers **convert points** into ERC-20 tokens.  
- **Fixed conversion rate** (e.g., 10 points = 1 token).  
- Workers can **claim tokens** anytime.  

### **4️⃣ Security Mechanisms**  
- **Ownable Modifier**: Only Task Manager can update worker points.  
- **ReentrancyGuard**: Prevents multiple reward claims in the same transaction.  
- **Mapping Storage**: Tracks worker points to avoid duplicate updates.  

---

## **🔹 Smart Contract Workflow**  

1️⃣ **Worker completes a task** (e.g., hashing, encryption).  
2️⃣ **Task Manager verifies result** and calls `updatePoints(worker, points)`.  
3️⃣ **Worker accumulates points** in the smart contract.  
4️⃣ **Worker calls `claimRewards()`** to convert points to tokens.  
5️⃣ **Tokens are transferred to the worker’s address**.  

---

## **🔹 Contract Functions**  

| **Function Name** | **Access** | **Purpose** |
|------------------|-----------|-------------|
| `updatePoints(address worker, uint256 points)` | **Only Owner** | Updates worker points after verification. |
| `getPoints(address worker)` | **Public** | Returns the worker's current points. |
| `claimRewards()` | **Public** | Converts points into ERC-20 tokens and transfers them. |
| `setRewardRate(uint256 rate)` | **Only Owner** | Changes the conversion rate of points to tokens. |

---

## **🔹 Contract Ownership & Security**  
- The contract **Owner** (Task Manager) manages rewards and settings.  
- **Ownership can be transferred** if needed.  
- **Only Owner** can modify key settings.  

---

## **🚀 Next Steps**  
✅ Deploy smart contract to **Ethereum Testnet**.  
✅ Implement **event logging** for transparency.  
✅ Add **DAO governance** to allow workers to vote on updates.  
