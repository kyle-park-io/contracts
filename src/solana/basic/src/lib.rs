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

    #[tokio::test]
    async fn test_solana_function() {
        let program_id = Pubkey::new_unique();

        let mut program_test = ProgramTest::default();

        program_test.add_program(
            "test_program",
            program_id,
            processor!(process_instruction), // ✅ add program
        );

        let (_banks_client, _payer, _recent_blockhash) = program_test.start().await;
    }
}
