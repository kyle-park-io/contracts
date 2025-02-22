use solana_program::{
    account_info::AccountInfo, entrypoint, entrypoint::ProgramResult, msg, pubkey::Pubkey,
};

entrypoint!(process_instruction);

fn process_instruction(
    _program_id: &Pubkey,
    _accounts: &[AccountInfo],
    _instruction_data: &[u8],
) -> ProgramResult {
    msg!("Solana program executed!");
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    use solana_program_test::{processor, ProgramTest};
    use solana_sdk::{
        pubkey::Pubkey,
        signer::{keypair::Keypair, Signer},
        transaction::Transaction,
    };

    #[tokio::test]
    async fn test_solana_function() {
        let program_id = Pubkey::new_unique();

        let mut program_test = ProgramTest::default();

        program_test.add_program(
            "test_program",
            program_id,
            processor!(process_instruction), // ✅ add program
        );

        let payer = Keypair::new();
        let (mut banks_client, payer, recent_blockhash) = program_test.start().await;

        let tx: Transaction = Transaction::new_signed_with_payer(
            &[], // No instructions
            Some(&payer.pubkey()),
            &[&payer],
            recent_blockhash,
        );
        assert!(tx.verify().is_ok());

        let result = banks_client.process_transaction(tx).await;
        assert!(result.is_ok());
    }
}
